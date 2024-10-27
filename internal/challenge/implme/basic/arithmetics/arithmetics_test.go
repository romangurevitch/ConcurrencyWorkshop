package arithmetics

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequentialSum(t *testing.T) {
	type args struct {
		ctx       context.Context
		inputSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "zero", args: args{ctx: nil, inputSize: 0}, want: 0},
		{name: "one", args: args{ctx: nil, inputSize: 1}, want: 1},
		{name: "two", args: args{ctx: nil, inputSize: 2}, want: 5},
		{name: "ten", args: args{ctx: nil, inputSize: 10}, want: 385},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SequentialSum(tt.args.inputSize), "SequentialSum(%v, %v)", tt.args.ctx, tt.args.inputSize)
		})
	}
}

func TestParallelSum(t *testing.T) {
	type args struct {
		ctx       context.Context
		inputSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "zero", args: args{ctx: nil, inputSize: 0}, want: 0},
		{name: "one", args: args{ctx: nil, inputSize: 1}, want: 1},
		{name: "two", args: args{ctx: nil, inputSize: 2}, want: 5},
		{name: "ten", args: args{ctx: nil, inputSize: 10}, want: 385},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ParallelSum(tt.args.inputSize), "ParallelSum(%v, %v)", tt.args.ctx, tt.args.inputSize)
		})
	}
}

// inputSizes defines the different input sizes for the benchmarks.
var inputSizes = []int{10, 100, 1000, 10000}

// BenchmarkSequentialSum runs the benchmark for the SequentialSum function with various input sizes.
func BenchmarkSequentialSum(b *testing.B) {
	for _, size := range inputSizes {
		b.Run("InputSize="+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SequentialSum(size)
			}
		})
	}
}

// BenchmarkParallelSum runs the benchmark for the ParallelSum function with various input sizes.
func BenchmarkParallelSum(b *testing.B) {
	for _, size := range inputSizes {
		b.Run("InputSize="+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ { //nolint
				ParallelSum(size)
			}
		})
	}
}
