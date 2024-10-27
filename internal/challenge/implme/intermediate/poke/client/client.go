package client

import (
	"strings"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

type Poke struct {
	Err      error
	Name     string
	ImageURL string
}

//go:generate  go run github.com/vektra/mockery/v2@v2.20.0 --with-expecter=true --name PokeClient
type PokeClient interface {
	FetchPokemon(id string) (*structs.Pokemon, error)
}

func New() PokeClient {
	return &pokeClient{}
}

type pokeClient struct {
}

func (p pokeClient) FetchPokemon(id string) (*structs.Pokemon, error) {
	pokemon, err := pokeapi.Pokemon(strings.TrimSpace(id))
	if err != nil {
		return nil, err
	}
	return &pokemon, nil
}
