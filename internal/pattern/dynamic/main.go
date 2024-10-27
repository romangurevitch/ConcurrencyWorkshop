package main

import (
	"context"
	"errors"
	"log/slog"
	"strconv"
	"sync"
	"time"

	"github.com/mtslzr/pokeapi-go"
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

var ErrNegativeValue = errors.New("negative value")

// FetchPokemonName just returns the Pokemon name as a string.
func FetchPokemonName(ctx context.Context, pokemonID int) (string, error) {
	pokemon, err := pokeapi.Pokemon(strconv.Itoa(pokemonID))
	if err != nil {
		return "", err
	}
	return pokemon.Name, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jobs := make(chan Job[int])
	limiter := rate.NewLimiter(rate.Every(100*time.Millisecond), 10) // Limit to 2 jobs per second with a burst of 10

	// Create a worker pool with 3 workers.
	results := NewRateLimited(ctx, limiter, jobs, FetchPokemonName)

	// This goroutine sends a new jobs.
	go func() {
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			default:
				jobs <- Job[int]{ID: i, Value: i}
			}
		}
	}()

	// Process the results.
	for result := range results {
		if result.Err != nil {
			slog.Error("Error processing job", "jobID", result.Job.ID, "error", result.Err)
			cancel()
			continue
		}
		slog.Info("Result for job", "jobID", result.Job.ID, "result", result.Value)
	}
}
