package main

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

// Result is a generic type to encapsulate the result of an operation.
type Result[T any] struct {
	Value T
	Err   error
}

// ProcessFunc defines a function type that processes a Result of type T and produces a Result of type U.
type ProcessFunc[T any, U any] func(context.Context, Result[T]) Result[U]

// Pipe reads Results of type T from inCh, processes them using the provided operation op,
// and sends the Results of type U on a new channel.
func Pipe[T any, U any](ctx context.Context, inCh <-chan Result[T], processFunc ProcessFunc[T, U]) <-chan Result[U] {
	outCh := make(chan Result[U])
	go func() {
		defer close(outCh) // Ensure the channel is closed when the goroutine exits.
		for {
			select {
			case <-ctx.Done():
				slog.Info("shutting down goroutine", "reason", ctx.Err())
				return
			case in, ok := <-inCh:
				if !ok {
					return // jobs channel closed, exit worker
				}
				outCh <- processFunc(ctx, in) // Process the result using processFunc and send it on the output channel.
			}
		}
	}()
	return outCh
}

// fetchPokemon fetches Pokémon data for a given ID.
func fetchPokemon(_ context.Context, result Result[int]) Result[structs.Pokemon] {
	pokemon, err := pokeapi.Pokemon(strconv.Itoa(result.Value))
	if err != nil {
		return Result[structs.Pokemon]{Err: err}
	}
	return Result[structs.Pokemon]{Value: pokemon}
}

// printPokemonName processes fetched Pokémon data to extract and print the Pokémon's name.
func printPokemonName(_ context.Context, result Result[structs.Pokemon]) Result[bool] {
	if result.Err != nil {
		slog.Error("Error processing job", "error", result.Err)
		return Result[bool]{Err: result.Err}
	}
	slog.Info("Pokemon Name", "name", result.Value.Name)
	return Result[bool]{Value: true}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure all pipelines are closed if main exits early.

	// Define maximum number of Pokémon to fetch.
	maxPokemon := 5

	// Create the pipeline.
	inputCh := make(chan Result[int])
	fetchCh := Pipe(ctx, inputCh, fetchPokemon)
	processCh := Pipe(ctx, fetchCh, printPokemonName)

	go func() {
		for i := 1; i <= maxPokemon; i++ {
			inputCh <- Result[int]{Value: i}
		}
		close(inputCh)
	}()

	// Wait for the last stage to complete.
	for result := range processCh {
		if result.Err != nil {
			slog.Error("Error", "error", result.Err)
		}
	}
}
