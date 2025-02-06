package arithmetics

import (
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
	wg := sync.WaitGroup{}
	wg.Add(inputSize)
	for i := 1; i <= inputSize; i++ {
		i := i
		go func() {
			defer wg.Done()
			sum.Add(process(int64(i)))
		}()
	}
	wg.Wait()
	return int(sum.Load())

}

func process(num int64) int64 {
	time.Sleep(time.Millisecond) // simulate processing time
	return num * num
}
