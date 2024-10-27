package rapidio

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/implme/advanced/rapidio/plotter"
	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/implme/advanced/rapidio/simulator"
)

func TestSequential(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err := runDefaultSimulation(t, ctx, NewSequential(), "sequential.png")
	require.NoError(t, err)
}

func TestConcurrent(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err := runDefaultSimulation(t, ctx, NewConcurrent(), "concurrent.png")
	require.NoError(t, err)
}

func runDefaultSimulation(t *testing.T, ctx context.Context, rapidIO RapidIO, resultsFilename string) error {
	sm := simulator.NewSimulator(ctx, getDefaultSimulatorConfig())
	rapidIO.HandleEvents(ctx, sm.GetChannels())
	sm.PrintConfig()

	sm.Start()
	rapidIO.Wait()
	slog.Info("Finished reading events", "number of events", sm.EventCount())

	require.Len(t, rapidIO.Results(), int(sm.EventCount()))
	return plotter.Plot(rapidIO.Results(), resultsFilename)
}

func getDefaultSimulatorConfig() simulator.Config {
	return simulator.Config{
		NumberOfChannels: 100,
		MaxInterval:      1000 * time.Microsecond, // Target a max frequency of an event every 100ms.
		MinInterval:      500 * time.Microsecond,  // Start at an event every 500ms.
		UpdateRate:       50 * time.Millisecond,
		IntervalStep:     2 * time.Microsecond,  // Increase frequency by 50ms every second.
		MaxJitter:        100 * time.Nanosecond, // Up to 10ms of jitter.
	}
}
