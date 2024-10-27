package rapidio

import (
	"context"

	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/implme/advanced/rapidio/simulator"
)

type concurrent struct {
}

func (c concurrent) HandleEvents(ctx context.Context, events []chan simulator.Event) {
	// TODO implement me
	panic("implement me")
}

func (c concurrent) Results() []simulator.EventResult {
	// TODO implement me
	panic("implement me")
}

func (c concurrent) Wait() {
	// TODO implement me
	panic("implement me")
}

// NewConcurrent constructs a new concurrent RapidIO instance.
func NewConcurrent() RapidIO {
	return &concurrent{}
}
