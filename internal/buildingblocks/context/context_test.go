package context

import (
	"context"
	"errors"
	"testing"
	"time"
)

// keyType is a type used for context value keys to avoid collisions.
type keyType string

// TestContextValue demonstrates the use of context values for storing and retrieving data.
func TestContextValue(t *testing.T) {
	ctx := context.Background()
	key := keyType("key")
	ctx = context.WithValue(ctx, key, "value")

	value := ctx.Value(key)
	if value != "value" {
		t.Errorf("Expected value 'value', got %v", value)
	}
}

// TestContextWithCancel demonstrates the use of context cancellation.
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

// TestContextWithTimeout demonstrates the use of context timeout.
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

// TestContextWithDeadline demonstrates the use of context deadline.
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

// TestContextPropagation demonstrates the propagation of context cancellation through multiple layers.
func TestContextPropagation(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())

	// Simulate a chain of operations each passing the context to the next
	go func(ctx context.Context) {
		go func(ctx context.Context) {
			// Wait for context cancellation
			<-ctx.Done()
		}(ctx)
	}(ctx)

	time.Sleep(time.Second) // Simulate some processing time
	cancelFunc()            // Cancel the context

	select {
	case <-ctx.Done():
		// Expected case
	case <-time.After(time.Second * 2):
		t.Error("Context cancellation propagation took too long")
	}
}

// TestWithCancelCause demonstrates the use of context.WithCancelCause.
func TestWithCancelCause(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithCancelCause(parent)

	myError := errors.New("my error")
	cancel(myError)

	if err := ctx.Err(); !errors.Is(err, context.Canceled) {
		t.Errorf("Expected context.Canceled, got %v", err)
	}

	if cause := context.Cause(ctx); !errors.Is(cause, myError) {
		t.Errorf("Expected 'my error', got %v", cause)
	}
}
