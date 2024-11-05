package fanoutin

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ErrNegativeValue = errors.New("negative value")

// Example squareNonNegative function that squares non-negative integer.
func squareNonNegative(_ context.Context, value int) (int, error) {
	if value < 0 {
		return 0, ErrNegativeValue
	}
	return value * value, nil
}

func TestFanOut(t *testing.T) {
	type args[T any, U any] struct {
		jobs    []Job[T]
		process ProcessFunc[T, U]
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
				jobs: func() []Job[int] {
					var jobs []Job[int]
					for i := 1; i <= 10; i++ {
						jobs = append(jobs, Job[int]{ID: i, Value: i})
					}
					return jobs
				}(),
				process: squareNonNegative,
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
				jobs:    []Job[int]{{ID: 1, Value: -1}},
				process: squareNonNegative,
			},
			want: []Result[int, int]{
				{Job: Job[int]{ID: 1, Value: -1}, Err: ErrNegativeValue},
			},
		},
		{
			name: "Cancelled context",
			args: args[int, int]{
				jobs:    []Job[int]{{ID: 1, Value: -1}},
				process: squareNonNegative,
			},
			cancel: true,
			want:   []Result[int, int]{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel() // ensure resources are cleaned up

			if tt.cancel {
				cancel() // cancel the context before waiting for the result
			}

			got := FanOut(ctx, tt.args.jobs, tt.args.process)
			var gotResults []Result[int, int]
			for result := range got {
				gotResults = append(gotResults, result)
			}
			assert.ElementsMatch(t, tt.want, gotResults)
		})
	}
}
