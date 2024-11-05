package fanoutin

import (
	"context"
	"log/slog"
	"sync"
)

// Job holds information about each job.
type Job[T any] struct {
	ID    int
	Value T
}

// Result holds information about each result.
type Result[T any, U any] struct {
	Job   Job[T]
	Value U
	Err   error
}

// ProcessFunc defines a function type for processing a value of type T to produce a value of type U, in a context-aware manner.
type ProcessFunc[T any, U any] func(context.Context, T) (U, error)

// FanOut creates a pool of workers.
func FanOut[T any, U any](ctx context.Context, jobs []Job[T], processFunc ProcessFunc[T, U]) chan Result[T, U] {
	results := make(chan Result[T, U], len(jobs))
	var wg sync.WaitGroup

	// Launch a new worker for each job.
	go func() {
		defer func() {
			// Close the results channel once all workers are done.
			wg.Wait()
			close(results)
		}()

		for i, job := range jobs {
			select {
			case <-ctx.Done():
				slog.Info("shutting down goroutine", "reason", ctx.Err(), "total jobs", len(jobs), "finished jobs", i)
				return
			default:
				wg.Add(1) // Increment the counter whenever a new job is received.
				go func(job Job[T]) {
					defer wg.Done() // Decrement the counter when the goroutine completes.

					value, err := processFunc(ctx, job.Value)
					results <- Result[T, U]{Job: job, Value: value, Err: err}
				}(job)
			}
		}
	}()

	return results
}
