package workerpool

import (
	"context"
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

// Worker function processes jobs and produces results.
func worker[T any, U any](ctx context.Context, jobs <-chan Job[T], results chan<- Result[T, U], processFunc ProcessFunc[T, U]) {
	for {
		select {
		case <-ctx.Done():
			return // context cancelled, exit worker
		case job, ok := <-jobs:
			if !ok {
				return // jobs channel closed, exit worker
			}
			value, err := processFunc(ctx, job.Value)
			results <- Result[T, U]{Job: job, Value: value, Err: err}
		}
	}
}

// CreateWorkerPool creates a pool of workers.
func CreateWorkerPool[T any, U any](ctx context.Context, numWorkers int, jobs <-chan Job[T], results chan<- Result[T, U], process ProcessFunc[T, U]) {
	wg := sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, jobs, results, process)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
}
