package goroutine

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/romangurevitch/concurrencyworkshop/internal/goroutine/counter"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var white = "\033[97m"
var cursorBack = "\033[G\033[K"

// UnexpectedResult what did you expect?
// Goroutines?
func UnexpectedResult() int {
	basicCounter := counter.NewBasicCounter()

	go func() {
		for i := 0; i < 1000; i++ {
			basicCounter.Inc() //counter++
		}
	}()

	return basicCounter.Count()
}

// UnexpectedResultFix is it fixed?
// WaitGroup
func UnexpectedResultFix() int {
	basicCounter := counter.NewBasicCounter()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			basicCounter.Inc()
		}
	}()

	wg.Wait()
	return basicCounter.Count()
}

// LetsMakeASmallChange ohh no!
// Race condition detection
func LetsMakeASmallChange() int {
	basicCounter := counter.NewBasicCounter()
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			basicCounter.Inc()
		}()
	}

	wg.Wait()
	return basicCounter.Count()
}

// FinallySomethingWorksAsExpected but is it?
// Locks, mutex, rwmutex, atomic
func FinallySomethingWorksAsExpected() int {
	basicCounter := counter.NewBasicCounter()
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			lock.Lock()
			basicCounter.Inc()
			lock.Unlock()
		}()
	}

	wg.Wait()
	return basicCounter.Count()
}

// FinallySomethingWorksAsExpectedAtomicCounter but is it?
func FinallySomethingWorksAsExpectedAtomicCounter() int {
	atomicCounter := counter.NewAtomicCounter()
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomicCounter.Inc()
		}()
	}

	wg.Wait()
	return atomicCounter.Count()
}

// NonStoppingGoRoutine is that a good idea?
func NonStoppingGoRoutine() int {
	atomicCounter := counter.NewAtomicCounter()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			inlinePrint(atomicCounter.Inc())
		}
	}()

	wg.Wait()
	return atomicCounter.Count()
}

// NonStoppingGoRoutineWithShutdown is it good enough though?
// channels, signals
func NonStoppingGoRoutineWithShutdown() (int, bool) {
	atomicCounter := counter.NewAtomicCounter()
	gracefulShutdown := false

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer func() { gracefulShutdown = true }()

		for {
			inlinePrint(atomicCounter.Inc())
		}
	}()

	<-sigs
	return atomicCounter.Count(), gracefulShutdown
}

// NonStoppingGoRoutineCorrectShutdown yes?
func NonStoppingGoRoutineCorrectShutdown() (int, bool) {
	atomicCounter := counter.NewAtomicCounter()
	wg := sync.WaitGroup{}
	gracefulShutdown := false

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() { gracefulShutdown = true }()

		for {
			select {
			case <-sigs:
				return
			default:
				inlinePrint(atomicCounter.Inc())
			}

		}
	}()

	wg.Wait()
	return atomicCounter.Count(), gracefulShutdown
}

// NonStoppingGoRoutineContext use context
// Context
func NonStoppingGoRoutineContext(ctx context.Context) (int, bool) {
	atomicCounter := counter.NewAtomicCounter()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	gracefulShutdown := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() { gracefulShutdown = true }()
		for {
			select {
			case <-ctx.Done():
				slog.Info("shutting down goroutine", "reason", ctx.Err())
				return
			case reason := <-sigs:
				slog.Info("shutting down goroutine", "reason", reason)
				return
			default:
				inlinePrint(atomicCounter.Inc())
			}
		}
	}()

	wg.Wait()
	return atomicCounter.Count(), gracefulShutdown
}

// NonStoppingGoRoutineContextBetter use context
func NonStoppingGoRoutineContextBetter(ctx context.Context) (int, bool) {
	atomicCounter := counter.NewAtomicCounter()

	ctx, cancelFunc := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancelFunc()
	wg := sync.WaitGroup{}
	gracefulShutdown := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() { gracefulShutdown = true }()

		for {
			select {
			case <-ctx.Done():
				slog.Info("shutting down goroutine", "reason", ctx.Err())
				return
			default:
				inlinePrint(atomicCounter.Inc())
			}
		}
	}()

	wg.Wait()
	return atomicCounter.Count(), gracefulShutdown
}

// NonStoppingGoRoutineContextBonus use context with tiny change
func NonStoppingGoRoutineContextBonus(ctx context.Context) (int, bool) {
	atomicCounter := counter.NewAtomicCounter()

	ctx, cancelFunc := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancelFunc()
	wg := sync.WaitGroup{}
	gracefulShutdown := false

	wg.Add(1)
	go func() {
		defer func() { gracefulShutdown = true }()
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				slog.Info("shutting down goroutine", "reason", ctx.Err())
				return
			default:
				inlinePrint(atomicCounter.Inc())
			}
		}
	}()

	wg.Wait()
	return atomicCounter.Count(), gracefulShutdown
}

func inlinePrint(result int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Print(yellow, cursorBack, result, reset)
}
