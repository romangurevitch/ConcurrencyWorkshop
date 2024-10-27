package main

import (
	"context"
	"log/slog"
	"strconv"
	"sync"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

type Result[T any] struct {
	Value T
	Err   error
}

type PubSub[T any] struct {
	subscribers sync.Map // key: topic (string), value: []chan Result[T]
}

func NewPubSub[T any]() *PubSub[T] {
	return &PubSub[T]{}
}

func (ps *PubSub[T]) Subscribe(topic string, ch chan Result[T]) {
	subscribers, _ := ps.subscribers.LoadOrStore(topic, []chan Result[T]{})
	// Append to the existing slice of channels
	ps.subscribers.Store(topic, append(subscribers.([]chan Result[T]), ch))
}

func (ps *PubSub[T]) Unsubscribe(topic string, ch chan Result[T]) {
	value, ok := ps.subscribers.Load(topic)
	if !ok {
		return // no subscribers for this topic
	}
	subscribers, ok := value.([]chan Result[T])
	if !ok {
		return
	}
	for i, subscriber := range subscribers {
		if subscriber == ch {
			// Remove the subscriber from the slice
			ps.subscribers.Store(topic, append(subscribers[:i], subscribers[i+1:]...))
			return
		}
	}
}

func (ps *PubSub[T]) Publish(topic string, message T) {
	value, ok := ps.subscribers.Load(topic)
	if !ok {
		return // no subscribers for this topic
	}
	subscribers, ok := value.([]chan Result[T])
	if !ok {
		return
	}
	for _, ch := range subscribers {
		select {
		case ch <- Result[T]{Value: message}:
		default: // if the channel is not ready to receive, move on to the next subscriber
		}
	}
}

// fetchPokemon fetches Pokémon data for a given ID.
func fetchPokemon(_ context.Context, pokeID int) (structs.Pokemon, error) {
	return pokeapi.Pokemon(strconv.Itoa(pokeID))
}

func main() {
	pubSub := NewPubSub[structs.Pokemon]()
	topicName := "pokemon"
	subscriber1 := make(chan Result[structs.Pokemon], 1)
	subscriber2 := make(chan Result[structs.Pokemon], 1)

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
