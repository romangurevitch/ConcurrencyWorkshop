package fixme

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"testing"
	"time"

	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/test"
)

// nolint
func TestNilChannel(t *testing.T) {
	test.ExitAfter(time.Millisecond)

	//var ch chan int --- Caution - should not use var
	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

	for val := range ch {
		slog.Info("successfully received", "value", val)
	}
}

// nolint
func TestClosedChannelWithoutOkCheck(t *testing.T) {
	test.ExitAfter(time.Millisecond)
	ch := make(chan int)

	go func() {
		ch <- 42
		close(ch)
	}()

	for {
		select {
		case val, ok := <-ch:
			if ok {
				slog.Info("received", "value", val)
			} else {
				slog.Info("channel closed")
				return
			}
			//caution -- slog.Info("received", "value", val)
		}
	}
}

// nolint
func TestClosedChannelWrite(t *testing.T) {
	defer test.ExpectNoPanic(t)

	ch := make(chan int, 1)
	defer close(ch) // caution - not included defer so writing will be panic on next line
	ch <- 5
}

// nolint
func TestUnlockingUnlockedLock(t *testing.T) {
	var mu sync.Mutex
	mu.Lock()
	mu.Unlock()
}

// nolint
func TestWaitGroupNegativeCounter(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		//wg.Done() --- caution --- two done statements for wg
	}()

	wg.Wait()
}

// nolint
func TestContextUsingPrimitivesAsKeys(t *testing.T) {
	type ctxKey string
	const key ctxKey = "myKey"
	ctx := context.WithValue(context.Background(), "myKey", "value1")

	if val, ok := ctx.Value(key).(string); !ok || val != "value1" {
		t.Fatalf("expected context to have 'value1' for 'myKey', got: %v", val)
	}
}

// nolint
func TestContextWithCancel(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Second * 2)
		cancelFunc() // Cancel the context after a delay
	}()

	select {
	case <-ctx.Done():
		if err := ctx.Err(); !errors.Is(err, context.Canceled) {
			t.Errorf("Expected context.Canceled, got %v", err)
		}
	case <-time.After(time.Second * 3):
		t.Error("Context cancellation took too long")
	}
}

// nolint
func TestContextWithTimeout(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc() // It's a good practice to call the cancel function even if the context times out

	select {
	case <-ctx.Done():
		if err := ctx.Err(); !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("Expected context.DeadlineExceeded, got %v", err)
		}
	case <-time.After(time.Second * 3):
		t.Error("Context timeout took too long")
	}
}

// nolint
func TestContextWithDeadline(t *testing.T) {
	deadline := time.Now().Add(time.Second * 2)
	ctx, cancelFunc := context.WithDeadline(context.Background(), deadline)
	defer cancelFunc() // It's a good practice to call the cancel function even if the context times out

	select {
	case <-ctx.Done():
		if err := ctx.Err(); !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("Expected context.DeadlineExceeded, got %v", err)
		}
	case <-time.After(time.Second * 3):
		t.Error("Context deadline took too long")
	}
}
