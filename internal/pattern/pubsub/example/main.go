package main

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"

	"github.com/romangurevitch/concurrencyworkshop/internal/pattern/pubsub"
)

// fetchPokemon fetches Pokémon data for a given ID.
func fetchPokemon(_ context.Context, pokeID int) (structs.Pokemon, error) {
	return pokeapi.Pokemon(strconv.Itoa(pokeID))
}

func main() {
	pubSub := pubsub.NewPubSub[structs.Pokemon]()
	topicName := "pokemon"
	subscriber1 := make(chan pubsub.Result[structs.Pokemon], 1)
	subscriber2 := make(chan pubsub.Result[structs.Pokemon], 1)

	pubSub.Subscribe(topicName, subscriber1)
	pubSub.Subscribe(topicName, subscriber2)

	poke, err := fetchPokemon(context.Background(), 1)
	if err != nil {
		slog.Error("Error fetching Pokemon", "error", err)
	}
	pubSub.Publish(topicName, poke)

	slog.Info("Received message on subscriber 1", "topic", topicName, "pokemon name", (<-subscriber1).Value.Name)
	slog.Info("Received message on subscriber 2", "topic", topicName, "pokemon name", (<-subscriber2).Value.Name)
}
