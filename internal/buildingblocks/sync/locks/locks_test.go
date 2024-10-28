package locks

import (
	"fmt"
	"sync"
	"testing"
)

// TestBasicLock demonstrates intermediate usage of sync.Mutex.
func TestBasicLock(t *testing.T) {
	var mu sync.Mutex
	var counter int
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	if counter != 10 {
		t.Errorf("Expected 2, got %v", counter)
	}
}

// TestDeferredUnlock demonstrates a common pattern.go of deferring the Unlock call.
func TestDeferredUnlock(t *testing.T) {
	var mu sync.Mutex
	var counter int
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			counter++
		}()
	}

	wg.Wait()
	if counter != 10 {
		t.Errorf("Expected 2, got %v", counter)
	}
}

func TestSemaphore(t *testing.T) {
	limit := 4
	// The semaphore channel
	sem := make(chan struct{}, limit)

	for i := 0; i < 10; i++ {
		sem <- struct{}{}
		go func(number int) {
			fmt.Println(number)
			<-sem
		}(i)
	}

	for i := 0; i < limit; i++ {
		sem <- struct{}{}
	}
}

// TestBadUsageLock demonstrates a bad usage of sync.Mutex.
func TestBadUsageLock(t *testing.T) {
	t.Skip("Comment out to demonstrate incorrect usage")

	var mu sync.Mutex
	var counter int
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Incorrect: forgetting to unlock the mutex will cause a deadlock
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			// missing mu.Unlock()
		}()
	}

	wg.Wait()
	if counter != 2 {
		t.Errorf("Expected 2, got %v", counter)
	}
}
