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
		val, _ := process(context.Background(), i)
		sum += val
	}
	return sum
}

// ParallelSum implement this method.
func ParallelSum(inputSize int) int {
	// panic("implement me!")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numOfJobs := inputSize
	var jobs []fanoutin.Job[int]
	for i := 1; i <= numOfJobs; i++ {
		jobs = append(jobs, fanoutin.Job[int]{ID: i, Value: i})
	}

	// Fan out
	results := fanoutin.FanOut(ctx, jobs, process)

	// Fan in
	sum := 0
	for result := range results {
		if result.Err != nil {
			slog.Error("Error processing job", "jobID", result.Job.ID, "error", result.Err)
			cancel()
			continue
		}
		slog.Info("Result for job", "jobID", result.Job.ID, "result", result.Value)
		sum += result.Value
	}
	return sum
}

func process(_ context.Context, num int) (int, error) {
	time.Sleep(time.Millisecond) // simulate processing time
	return num * num, nil
}
