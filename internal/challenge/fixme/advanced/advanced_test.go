package fixme

import (
	"context"
	"errors"
	"log/slog"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/test"
)

// nolint
func TestWaitGroupWithoutDefer(t *testing.T) {
	test.ExitAfter(100 * time.Millisecond)

	wg := sync.WaitGroup{}
	finishedSuccessfully := false

	finishedFunc := func() {
		finishedSuccessfully = true
		runtime.Goexit()
	}

	wg.Add(1)
	go func() {
		defer finishedFunc()
		wg.Done()
	}()

	wg.Wait()
	require.True(t, finishedSuccessfully)
}

// nolint
func TestErrGroupWithoutWithContext(t *testing.T) {
	test.ExitAfter(10 * time.Millisecond)
	expectedErr := errors.New("error")
	ctx, cancelFunc := context.WithCancelCause(context.Background())
	group := errgroup.Group{}

	group.Go(func() error {
		cancelFunc(expectedErr)
		return expectedErr
	})

	group.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	if err := group.Wait(); err != nil {
		require.ErrorIs(t, err, expectedErr)
	}
}

// nolint
func TestContextIgnoringCancellation(t *testing.T) {
	test.ExitAfter(10 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	inputCh := make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		// Waiting on input
		case <-inputCh:
		case <-ctx.Done():
		}
	}()

	wg.Wait()
}

// nolint
func TestMultipleProducersCloseChannel(t *testing.T) {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	producer := func() {
		defer wg.Done()
		ch <- 1
	}

	wg.Add(2)
	go producer()
	go producer()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		slog.Info("successfully received", "value", val)
	}
}
