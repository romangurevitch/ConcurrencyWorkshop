package locks

import (
	"sync"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	lock := sync.Mutex{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lock.Lock()
		lock.Unlock() // nolint
	}
}

func BenchmarkRWMutex_Lock(b *testing.B) {
	lock := sync.RWMutex{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lock.Lock()
		lock.Unlock() // nolint
	}
}

func BenchmarkRWMutex_RLock(b *testing.B) {
	lock := sync.RWMutex{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lock.RLock()
		lock.RUnlock() // nolint
	}
}

func BenchmarkWaitGroup_Add(b *testing.B) {
	wg := sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
	}
}

func BenchmarkWaitGroup_Done(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Done()
	}
}
