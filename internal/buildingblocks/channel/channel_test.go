package channel

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/test"
)

// TestBasicChannelUsage demonstrates intermediate usage of channels for communication between goroutines.
func TestBasicChannelUsage(t *testing.T) {
	ch := make(chan int) // Create a new channel

	go func() {
		ch <- 42 // Send a value into the channel
	}()

	value := <-ch // Receive a value from the channel
	if value != 42 {
		t.Errorf("Expected 42, got %d", value)
	}
}

// TestBufferedChannel demonstrates the use of buffered channels.
func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 2) // Create a buffered channel with a capacity of 2

	ch <- 42 // Send values into the channel
	ch <- 43

	// Receive values from the channel and check them
	if v := <-ch; v != 42 {
		t.Errorf("Expected 42, got %d", v)
	}
	if v := <-ch; v != 43 {
		t.Errorf("Expected 43, got %d", v)
	}
}

// TestCloseChannel demonstrates how to close a channel and how to read from a closed channel using a range loop.
func TestCloseChannel(t *testing.T) {
	ch := make(chan int, 2) // Create a buffered channel

	ch <- 42 // Send values into the channel
	ch <- 43
	close(ch) // Close the channel

	// Range over the channel to receive values until the channel is closed
	for v := range ch {
		fmt.Println(v)
	}
}

// TestSendOnClosedChannel checks for a panic when sending to a closed channel.
func TestSendOnClosedChannel(t *testing.T) {
	ch := make(chan int) // Create a new channel
	close(ch)            // Close the channel

	defer test.ExpectPanic(t)
	ch <- 42 // Attempt to send to a closed channel
}

// TestCloseClosedChannel checks for a panic when closing a closed channel.
func TestCloseClosedChannel(t *testing.T) {
	ch := make(chan int) // Create a new channel
	close(ch)            // Close the channel

	defer test.ExpectPanic(t)
	close(ch) // Attempt to close a closed channel
}

// TestSelectDefault demonstrates how to use a select statement with a default case to prevent blocking.
func TestSelectDefault(t *testing.T) {
	ch := make(chan int) // Create a new channel

	// Use a select statement to attempt to receive from the channel, with a default case to prevent blocking
	select {
	case v := <-ch:
		t.Errorf("Received unexpected value: %d", v)
	default:
		// Expected case
	}
}

// TestProducerConsumer demonstrates an intermediate producer-consumer pattern.go using a channel.
func TestProducerConsumer(t *testing.T) {
	ch := make(chan int) // Create a new channel

	// Producer goroutine: sends values into the channel
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // Close the channel when done
	}()

	// Consumer: receive values from the channel until it's closed
	for v := range ch {
		fmt.Println(v)
	}
}

// TestWithContext demonstrates how to use a context to control channel operations.
func TestWithContext(t *testing.T) {
	ch := make(chan int)                                                           // Create a new channel
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond) // Create a context with a timeout
	defer cancel()                                                                 // Ensure the context is canceled to release resources

	// Goroutine: attempts to send a value into the channel, but exits if the context expires first
	go func() {
		select {
		case <-ctx.Done():
			close(ch) // Close the channel if the context expires
		case ch <- 42:
			// Sent value
		}
	}()

	// Wait to receive a value from the channel or for the context to expire
	select {
	case v := <-ch:
		if v != 42 {
			t.Errorf("Expected 42, got %d", v)
		}
	case <-ctx.Done():
		t.Error("Context timeout expired before value was received")
	}
}
