package rapidio

import (
	"context"
	"time"

	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/implme/advanced/rapidio/simulator"
)

type RapidIO interface {
	HandleEvents(ctx context.Context, events []chan simulator.Event)
	Results() []simulator.EventResult
	Wait()
}

func EventHandler(event simulator.Event) simulator.EventResult {
	return simulator.EventResult{Event: event, HandledAt: time.Now()}
}
