package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"

	"github.com/romangurevitch/concurrencyworkshop/internal/pattern/future"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure all resources are cleaned up

	getPokeFuture := future.NewFuture(ctx, func(ctx context.Context) (structs.Pokemon, error) {
		return pokeapi.Pokemon("pikachu")
	})

	// Optionally, do some other work here while waiting for the getPokeFuture result...

	// Now wait for the result:
	result := getPokeFuture.Result()
	if result.Err != nil {
		slog.Error("Error fetching Pokémon details", "error", result.Err)
		return
	}
	slog.Info("Fetched Pokémon details", "pokemonName", result.Value.Name)
}
