package main

import (
	"context"
	"errors"
	"log/slog"

	"github.com/romangurevitch/concurrencyworkshop/internal/pattern/fanoutin"
)

var ErrNegativeValue = errors.New("negative value")

// Example squareNonNegative function that squares non-negative integer.
func squareNonNegative(_ context.Context, value int) (int, error) {
	if value < 0 {
		return 0, ErrNegativeValue
	}
	return value * value, nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numOfJobs := 10
	var jobs []fanoutin.Job[int]
	for i := 1; i <= numOfJobs; i++ {
		jobs = append(jobs, fanoutin.Job[int]{ID: i, Value: i})
	}

	// Fan out
	results := fanoutin.FanOut(ctx, jobs, squareNonNegative)

	// Fan in
	for result := range results {
		if result.Err != nil {
			slog.Error("Error processing job", "jobID", result.Job.ID, "error", result.Err)
			cancel()
			continue
		}
		slog.Info("Result for job", "jobID", result.Job.ID, "result", result.Value)
	}
}
