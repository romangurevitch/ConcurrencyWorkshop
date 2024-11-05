package main

import (
	"context"
	"log/slog"
	"strconv"
	"time"

	"github.com/mtslzr/pokeapi-go"
	"golang.org/x/time/rate"

	"github.com/romangurevitch/concurrencyworkshop/internal/pattern/dynamic"
)

// FetchPokemonName just returns the Pokemon name as a string.
func FetchPokemonName(_ context.Context, pokemonID int) (string, error) {
	pokemon, err := pokeapi.Pokemon(strconv.Itoa(pokemonID))
	if err != nil {
		return "", err
	}
	return pokemon.Name, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jobs := make(chan dynamic.Job[int])
	limiter := rate.NewLimiter(rate.Every(100*time.Millisecond), 10) // Limit to 2 jobs per second with a burst of 10

	// Create a worker pool with 3 workers.
	results := dynamic.NewRateLimited(ctx, limiter, jobs, FetchPokemonName)

	// This goroutine sends a new jobs.
	go func() {
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			default:
				jobs <- dynamic.Job[int]{ID: i, Value: i}
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
