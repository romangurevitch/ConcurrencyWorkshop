package rapidio

import (
	"context"
	"log/slog"
	"sync"

	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/implme/advanced/rapidio/simulator"
)

// sequential is a type that implements the RapidIO interface sequentially.
// It collects results from a series of event channels into a slice.
type sequential struct {
	results []simulator.EventResult
	wg      *sync.WaitGroup
}

// NewSequential constructs a new sequential RapidIO instance.
func NewSequential() RapidIO {
	return &sequential{wg: &sync.WaitGroup{}}
}

// HandleEvents begins processing events from a slice of channels.
// It listens for event emissions on each channel and processes them sequentially.
func (s *sequential) HandleEvents(ctx context.Context, events []chan simulator.Event) {
	green := make([]chan simulator.Event, len(events))
	copy(green, events)

	var red []chan simulator.Event

	// Process all events in one goroutine sequentially
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			if len(green) == 0 {
				return // If there are no more channels to process, exit the goroutine.
			}

			for i := 0; i < len(green); i++ {
				select {
				case <-ctx.Done():
					slog.Error("RapidIO", "error", ctx.Err())
					return
				case event, ok := <-green[i]:
					if !ok {
						break
					}
					// Process the event and append the result to the 'results' slice.
					s.results = append(s.results, EventHandler(event))
					red = append(red, green[i])
				}
			}

			// Rotate the slices: move the processed channels from 'green' to 'red',
			// and reset 'red' for the next iteration.
			green = red
			red = nil
		}
	}()
}

// Wait blocks until all events have been handled.
func (s *sequential) Wait() {
	s.wg.Wait() // Wait for the event handling goroutine to finish.
}

// Results returns the slice of event results collected by the sequential handler.
func (s *sequential) Results() []simulator.EventResult {
	return s.results
}
