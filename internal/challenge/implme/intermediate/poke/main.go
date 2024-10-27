package main

import (
	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/implme/intermediate/poke/app"
	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/implme/intermediate/poke/client"
)

func main() {
	pokeAPP := app.NewPokeApp(client.New())
	pokeAPP.Start()
}
