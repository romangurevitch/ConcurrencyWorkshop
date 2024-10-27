package app

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/implme/intermediate/poke/client"
	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/implme/intermediate/poke/client/mocks"
)

func Test_pokeAPP_OnChangedNonBlocking(t *testing.T) {
	type fields struct {
		pokeClient func(t *testing.T) client.PokeClient
	}
	type args struct {
		ID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test non blocking",
			fields: fields{func(t *testing.T) client.PokeClient {
				pokeClient := mocks.NewPokeClient(t)
				pokeClient.EXPECT().FetchPokemon(mock.Anything).Maybe().Run(func(_ mock.Arguments) {
					// If the call is blocking the test will not finish successfully
					var nilChan chan struct{}
					nilChan <- struct{}{} //nolint
				})
				return pokeClient
			}},
			args: args{ID: "pikachu"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pokeAPP{
				pokeClient: tt.fields.pokeClient(t),
			}
			p.OnChangedNonBlocking(tt.args.ID)
		})
	}
}
