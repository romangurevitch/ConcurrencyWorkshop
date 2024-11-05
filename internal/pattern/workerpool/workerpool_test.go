package workerpool

import (
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var ErrNegativeValue = errors.New("negative value")

// Example squareNonNegative function that squares non-negative integer.
func squareNonNegative(ctx context.Context, value int) (int, error) {
	slog.Info("squareNonNegative", "value", value, "ctx err", ctx.Err())
	if value < 0 {
		return 0, ErrNegativeValue
	}
	return value * value, nil
}

func TestWorkerPool(t *testing.T) {
	type args[T any, U any] struct {
		numWorkers int
		jobs       []Job[T]
		process    ProcessFunc[T, U]
	}
	type testCase[T any, U any] struct {
		name   string
		args   args[T, U]
		cancel bool // whether to cancel the context before waiting for the result
		want   []Result[T, U]
	}

	tests := []testCase[int, int]{
		{
			name: "Basic Test",
			args: args[int, int]{
				numWorkers: 2,
				jobs: []Job[int]{
					{ID: 1, Value: 2},
					{ID: 2, Value: 3},
					{ID: 3, Value: -1}, // negative value, should result in error
				},
				process: squareNonNegative,
			},
			want: []Result[int, int]{
				{Job: Job[int]{ID: 1, Value: 2}, Value: 4},
				{Job: Job[int]{ID: 2, Value: 3}, Value: 9},
				{Job: Job[int]{ID: 3, Value: -1}, Err: ErrNegativeValue},
			},
		},
		{
			name: "No Jobs",
			args: args[int, int]{
				numWorkers: 2,
				jobs:       []Job[int]{}, // no jobs to squareNonNegative
				process:    squareNonNegative,
			},
			want: []Result[int, int]{},
		},
		{
			name: "Single Worker",
			args: args[int, int]{
				numWorkers: 1,
				jobs: []Job[int]{
					{ID: 1, Value: 2},
					{ID: 2, Value: 3},
				},
				process: squareNonNegative,
			},
			want: []Result[int, int]{
				{Job: Job[int]{ID: 1, Value: 2}, Value: 4},
				{Job: Job[int]{ID: 2, Value: 3}, Value: 9},
			},
		},
		{
			name: "Multiple Workers More Jobs",
			args: args[int, int]{
				numWorkers: 3,
				jobs: []Job[int]{
					{ID: 1, Value: 2},
					{ID: 2, Value: 3},
					{ID: 3, Value: 4},
					{ID: 4, Value: 5},
				},
				process: squareNonNegative,
			},
			want: []Result[int, int]{
				{Job: Job[int]{ID: 1, Value: 2}, Value: 4},
				{Job: Job[int]{ID: 2, Value: 3}, Value: 9},
				{Job: Job[int]{ID: 3, Value: 4}, Value: 16},
				{Job: Job[int]{ID: 4, Value: 5}, Value: 25},
			},
		},
		{
			name: "Cancelled context",
			args: args[int, int]{
				numWorkers: 3,
				jobs: []Job[int]{
					{ID: 1, Value: 2},
					{ID: 2, Value: 3},
					{ID: 3, Value: 4},
					{ID: 4, Value: 5},
				},
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

			jobsChan := make(chan Job[int], len(tt.args.jobs))
			resultsChan := make(chan Result[int, int], len(tt.args.jobs))
			CreateWorkerPool(ctx, tt.args.numWorkers, jobsChan, resultsChan, tt.args.process)
			if tt.cancel {
				cancel()                           // cancel the context before waiting for the result
				time.Sleep(100 * time.Millisecond) // propagation of the context takes time
			}

			for _, job := range tt.args.jobs {
				jobsChan <- job
			}
			close(jobsChan) // close jobs channel after feeding all jobs

			var gotResults []Result[int, int]
			for result := range resultsChan {
				gotResults = append(gotResults, result)
			}

			assert.ElementsMatch(t, tt.want, gotResults)
		})
	}
}
