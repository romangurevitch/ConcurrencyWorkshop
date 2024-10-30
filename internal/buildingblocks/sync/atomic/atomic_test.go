package atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

// TestBasicAtomic demonstrates intermediate usage of atomic.Int64.
func TestBasicAtomic(t *testing.T) {
	counter := atomic.Int64{}
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		counter.Add(1)
	}()

	go func() {
		defer wg.Done()
		counter.Add(1)
	}()

	wg.Wait()
	if counter.Load() != 2 {
		t.Errorf("Expected 2, got %v", counter.Load())
	}
}

// TestBadUsageAtomic demonstrates a bad usage of atomic.Int64.
func TestBadUsageAtomic(t *testing.T) {
	t.Skip("Comment out to demonstrate incorrect usage")
	var counter atomic.Int64
	wg := sync.WaitGroup{}

	// Incorrect: Updating the counter without atomic operations
	iterations := 10000
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Store(counter.Load() + 1)
		}()
	}

	wg.Wait()
	if counter.Load() == int64(iterations) {
		t.Errorf("Expected race condition that results in less than %d, got %v", iterations, counter.Load())
	}
}

// TestMixedUsageAtomic demonstrates the goroutine of mixing atomic and non-atomic operations.
// Run with and without race detection flag.
func TestMixedUsageAtomic(t *testing.T) {
	var counter int64
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		atomic.AddInt64(&counter, 1)
	}()

	// Incorrect: Mixing atomic and non-atomic operations
	go func() {
		defer wg.Done()
		counter++
	}()

	wg.Wait()
	if atomic.LoadInt64(&counter) != 2 {
		t.Errorf("Expected 2, got %v", counter)
	}
}
