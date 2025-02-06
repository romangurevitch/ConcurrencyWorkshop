package arithmetics

import (
	"context"
	"log/slog"
	"time"

	"github.com/romangurevitch/concurrencyworkshop/internal/pattern/fanoutin"
)

func SequentialSum(inputSize int) int {
	sum := 0
	for i := 1; i <= inputSize; i++ {
		val, err := process(context.Background(), i)
		if err == nil {
			sum += val
		}
	}
	return sum
}

func ParallelSum(inputSize int) int {
	sum := 0
	ctx, cancel := context.WithCancel(context.Background())

	var jobs []fanoutin.Job[int]
	for i := 1; i <= inputSize; i++ {
		jobs = append(jobs, fanoutin.Job[int]{ID: i, Value: i})
	}

	// Fan out
	results := fanoutin.FanOut(ctx, jobs, process)

	// Fan in
	for result := range results {
		if result.Err != nil {
			slog.Error("Error processing job", "jobID", result.Job.ID, "error", result.Err)
			cancel()
			continue
		}
		sum += result.Value
	}
	return sum
}

func process(_ context.Context, num int) (int, error) {
	time.Sleep(time.Millisecond) // simulate processing time
	return num * num, nil
}
