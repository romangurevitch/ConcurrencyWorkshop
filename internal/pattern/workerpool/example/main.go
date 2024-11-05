package main

import (
	"context"
	"log/slog"
	"strconv"
	"time"

	"github.com/mtslzr/pokeapi-go"

	"github.com/romangurevitch/concurrencyworkshop/internal/pattern/workerpool"
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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	jobs := make(chan workerpool.Job[int])
	results := make(chan workerpool.Result[int, string])

	// Create a worker pool with 3 workers.
	workerpool.CreateWorkerPool(ctx, 3, jobs, results, FetchPokemonName)

	// This goroutine sends a new job every second.
	go func() {
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			default:
				jobs <- workerpool.Job[int]{ID: i, Value: i}
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
