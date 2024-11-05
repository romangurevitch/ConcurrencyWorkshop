package main

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"

	"github.com/romangurevitch/concurrencyworkshop/internal/pattern/pipeline"
)

// fetchPokemon fetches Pokémon data for a given ID.
func fetchPokemon(_ context.Context, result pipeline.Result[int]) pipeline.Result[structs.Pokemon] {
	pokemon, err := pokeapi.Pokemon(strconv.Itoa(result.Value))
	if err != nil {
		return pipeline.Result[structs.Pokemon]{Err: err}
	}
	return pipeline.Result[structs.Pokemon]{Value: pokemon}
}

// printPokemonName processes fetched Pokémon data to extract and print the Pokémon's name.
func printPokemonName(_ context.Context, result pipeline.Result[structs.Pokemon]) pipeline.Result[bool] {
	if result.Err != nil {
		slog.Error("Error processing job", "error", result.Err)
		return pipeline.Result[bool]{Err: result.Err}
	}
	slog.Info("Pokemon Name", "name", result.Value.Name)
	return pipeline.Result[bool]{Value: true}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure all pipelines are closed if main exits early.

	// Define maximum number of Pokémon to fetch.
	maxPokemon := 5

	// Create the pipeline.
	inputCh := make(chan pipeline.Result[int])
	fetchCh := pipeline.Pipe(ctx, inputCh, fetchPokemon)
	processCh := pipeline.Pipe(ctx, fetchCh, printPokemonName)

	go func() {
		for i := 1; i <= maxPokemon; i++ {
			inputCh <- pipeline.Result[int]{Value: i}
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
