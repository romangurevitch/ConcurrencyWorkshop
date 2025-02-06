package arithmetics

import (
	"sync"
	"sync/atomic"
	"time"
)

func SequentialSum(inputSize int) int {
	sum := 0
	for i := 1; i <= inputSize; i++ {
		sum += process(i)
	}
	return sum
}

// ParallelSum implement this method.
func ParallelSum(inputSize int) int {
	sum := atomic.Int32{}
	wg := sync.WaitGroup{}
	for i := 1; i <= inputSize; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sum.Add(int32(process(num)))
		}(i)
	}
	wg.Wait()
	return int(sum.Load())
}

func process(num int) int {
	time.Sleep(time.Millisecond) // simulate processing time
	return num * num
}
