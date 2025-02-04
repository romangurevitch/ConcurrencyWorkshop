package arithmetics

import (
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
func ParallelSum(_ int) int {
	panic("implement me!")
}

func process(num int) int {
	time.Sleep(time.Millisecond) // simulate processing time
	return num * num
}
