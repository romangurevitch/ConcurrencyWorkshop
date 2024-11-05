package dynamic

import (
	"context"
	"log/slog"
	"sync"

	"golang.org/x/time/rate"
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

// NewRateLimited creates a rate-limited worker pool.
func NewRateLimited[T any, U any](ctx context.Context, limiter *rate.Limiter, jobs <-chan Job[T], processFunc ProcessFunc[T, U]) <-chan Result[T, U] {
	results := make(chan Result[T, U], limiter.Burst())

	go func() {
		wg := sync.WaitGroup{}
		defer func() {
			// Close the results channel once all workers are done.
			wg.Wait()
			close(results)
		}()

		for {
			select {
			case <-ctx.Done():
				slog.Info("shutting down goroutine", "reason", ctx.Err())
				return
			case job, ok := <-jobs:
				if !ok {
					return // jobs channel closed, exit worker
				}
				if err := limiter.Wait(context.Background()); err != nil { // context shutdown is handled elsewhere.
					results <- Result[T, U]{Job: job, Err: err}
					return
				}
				wg.Add(1)
				go func(job Job[T]) {
					defer wg.Done()
					value, err := processFunc(ctx, job.Value)
					results <- Result[T, U]{Job: job, Value: value, Err: err}
				}(job)
			}
		}
	}()

	return results
}
