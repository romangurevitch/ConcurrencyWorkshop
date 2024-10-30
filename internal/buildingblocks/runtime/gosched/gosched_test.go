package gosched

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// TestBasicGosched demonstrates intermediate usage of runtime.Gosched.
func TestBasicGosched(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1")
		runtime.Gosched()
		fmt.Println("Goroutine 1 Resumed")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2")
	}()

	wg.Wait()
}

// TestGoschedWithLongComputation demonstrates the usage of runtime.Gosched with long computation tasks.
func TestGoschedWithLongComputation(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1e6; i++ {
			if i%1000 == 0 {
				runtime.Gosched() // Allow other goroutines to run.
			}
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine with less work")
	}()

	wg.Wait()
}

// TestBadUsageGosched demonstrates a bad usage of runtime.Gosched.
func TestBadUsageGosched(t *testing.T) {
	t.Skip("Comment out to demonstrate Gosched incorrect usage")
	for {
		runtime.Gosched() // Incorrect: This creates a busy loop, which will keep the CPU busy without doing useful work.
	}
}
