package atomic

import (
	"sync/atomic"
	"testing"
)

func BenchmarkAtomic_atomicInt64_Add(b *testing.B) {
	atomicInt64 := atomic.Int64{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomicInt64.Add(1)
	}
}

func BenchmarkAtomic_atomicInt64_CompareAndSwap(b *testing.B) {
	atomicInt64 := atomic.Int64{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomicInt64.CompareAndSwap(1, 1)
	}
}

func BenchmarkAtomic_atomicInt64_Load(b *testing.B) {
	atomicInt64 := atomic.Int64{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomicInt64.Load()
	}
}

func BenchmarkAtomic_Int64_Add(b *testing.B) {
	value := int64(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&value, 1)
	}
}

func BenchmarkAtomic_Int64_CompareAndSwap(b *testing.B) {
	value := int64(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomic.CompareAndSwapInt64(&value, 1, 1)
	}
}

func BenchmarkAtomic_Int64_Load(b *testing.B) {
	value := int64(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomic.LoadInt64(&value)
	}
}
