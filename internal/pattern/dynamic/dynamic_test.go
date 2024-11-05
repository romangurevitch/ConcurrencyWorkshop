package dynamic

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

var ErrNegativeValue = errors.New("negative value")

func TestNewRateLimited(t *testing.T) {
	type args[T any, U any] struct {
		limiter     *rate.Limiter
		jobs        []Job[T]
		processFunc ProcessFunc[T, U]
	}
	type testCase[T any, U any] struct {
		name   string
		args   args[T, U]
		cancel bool // whether to cancel the context before waiting for the result
		want   []Result[T, U]
	}

	tests := []testCase[int, int]{
		{
			name: "Positive Values",
			args: args[int, int]{
				limiter: rate.NewLimiter(rate.Every(time.Millisecond*100), 10),
				jobs: func() []Job[int] {
					var jobs []Job[int]
					for i := 1; i <= 10; i++ {
						jobs = append(jobs, Job[int]{ID: i, Value: i})
					}
					return jobs
				}(),
				processFunc: squareNonNegative,
			},
			want: func() []Result[int, int] {
				var results []Result[int, int]
				for i := 1; i <= 10; i++ {
					results = append(results, Result[int, int]{Job: Job[int]{ID: i, Value: i}, Value: i * i})
				}
				return results
			}(),
		},
		{
			name: "Negative Value",
			args: args[int, int]{
				limiter: rate.NewLimiter(rate.Every(time.Millisecond*100), 10),
				jobs: []Job[int]{
					{ID: 1, Value: -1},
				},
				processFunc: squareNonNegative,
			},
			want: []Result[int, int]{
				{Job: Job[int]{ID: 1, Value: -1}, Err: ErrNegativeValue},
			},
		},
		{
			name: "Cancelled context",
			args: args[int, int]{
				limiter: rate.NewLimiter(rate.Every(time.Millisecond*100), 10),
				jobs: []Job[int]{
					{ID: 1, Value: -1},
				},
				processFunc: squareNonNegative,
			},
			cancel: true,
			want:   []Result[int, int]{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel() // ensure resources are cleaned up

			jobsChan := make(chan Job[int], len(tt.args.jobs))
			resultsChan := NewRateLimited(ctx, tt.args.limiter, jobsChan, tt.args.processFunc)
			if tt.cancel {
				cancel()                           // cancel the context before waiting for the result
				time.Sleep(100 * time.Millisecond) // propagation of the context takes time
			}

			for _, job := range tt.args.jobs {
				jobsChan <- job
			}
			close(jobsChan)

			var gotResults []Result[int, int]
			for result := range resultsChan {
				gotResults = append(gotResults, result)
			}
			assert.ElementsMatch(t, tt.want, gotResults)
		})
	}
}

// Example squareNonNegative function that squares non-negative integer.
func squareNonNegative(_ context.Context, value int) (int, error) {
	if value < 0 {
		return 0, ErrNegativeValue
	}
	return value * value, nil
}
