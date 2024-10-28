package goexit

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// TestBasicGoexit demonstrates intermediate usage of runtime.Goexit.
func TestBasicGoexit(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Deferred function in goroutine")

		fmt.Println("About to exit goroutine")
		runtime.Goexit()

		// This line will not be reached
		t.Error("This line should not be reached")
	}()

	wg.Wait()
}

// TestGoexitInMainGoroutine demonstrates a bad usage of runtime.Goexit.
func TestGoexitInMainGoroutine(t *testing.T) {
	t.Skip("Comment out to demonstrate Goexit incorrect usage")

	defer fmt.Println("Deferred function in TestGoexitInMainGoroutine")

	// Incorrect: Calling runtime.Goexit in the pipeline goroutine
	// will cause the program to exit before anything else can happen.
	// Unlike panics, runtime.Goexit() does not interact with recover(),
	// and there's no mechanism to stop or reverse the termination of the goroutine.
	runtime.Goexit()

	t.Error("This line should not be reached")
}

// TestGoexitWithoutDefer demonstrates a pitfall where defers are not used with runtime.Goexit.
func TestGoexitWithoutDefer(t *testing.T) {
	t.Skip("Comment out to demonstrate Goexit incorrect usage")

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		// No defer to call wg.Done, this will cause a deadlock
		fmt.Println("About to exit goroutine")
		runtime.Goexit()
		wg.Done()
	}()

	wg.Wait() // This will deadlock
}
