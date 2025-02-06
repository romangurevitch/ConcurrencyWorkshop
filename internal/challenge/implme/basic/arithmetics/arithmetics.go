package arithmetics

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

func SequentialSum(inputSize int) int {
	sum := 0
	for i := 1; i <= int(inputSize); i++ {
		sum += int(process(int64(i)))
	}
	return sum
}

// ParallelSum implement this method.
func ParallelSum(inputSize int) int {
	var sum atomic.Int64
	ctx := context.Background()
	wg := sync.WaitGroup{}
	wg.Add(inputSize)
	for i := 1; i <= inputSize; i++ {
		i := i
		go func(ctx context.Context) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				sum.Add(process(int64(i)))
			}
		}(ctx)
	}
	wg.Wait()
	return int(sum.Load())

}

func process(num int64) int64 {
	time.Sleep(time.Millisecond) // simulate processing time
	return num * num
}
