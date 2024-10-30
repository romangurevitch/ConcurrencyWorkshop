package pool

import (
	"sync"
	"testing"
)

func BenchmarkPool_Put(b *testing.B) {
	pool := sync.Pool{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Put(&struct{}{})
	}
}

func BenchmarkPool_Get(b *testing.B) {
	pool := sync.Pool{}
	for i := 0; i < b.N; i++ {
		pool.Put(&struct{}{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}
