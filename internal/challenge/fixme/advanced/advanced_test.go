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
	cancelFn := test.ExitWithCancelAfter(context.Background(), time.Second)
	defer cancelFn()

	wg := sync.WaitGroup{}
	finishedSuccessfully := false

	finishedFunc := func() {
		finishedSuccessfully = true
		runtime.Goexit()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		finishedFunc()
	}()

	wg.Wait()
	require.True(t, finishedSuccessfully)
}

// nolint
func TestErrGroupWithoutWithContext(t *testing.T) {
	cancelFn := test.ExitWithCancelAfter(context.Background(), time.Second)
	defer cancelFn()

	expectedErr := errors.New("error")
	ctx := context.Background()
	group := errgroup.Group{}

	group.Go(func() error {
		return expectedErr
	})

	go func() {
		err := func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			}
		}()
		if err != nil {
			require.ErrorIs(t, err, context.Canceled)
		}
	}()

	if err := group.Wait(); err != nil {
		require.ErrorIs(t, err, expectedErr)
	}
}

// nolint
func TestContextIgnoringCancellation(t *testing.T) {
	cancelFn := test.ExitWithCancelAfter(context.Background(), time.Second)
	defer cancelFn()

	_, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	inputCh := make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		// Waiting on input
		case <-inputCh:
		}
	}()

	inputCh <- true

	wg.Wait()
}

// nolint
func TestMultipleProducersCloseChannel(t *testing.T) {
	ch := make(chan int, 2)
	wg := sync.WaitGroup{}

	producer := func() {
		defer wg.Done()
		ch <- 1
	}

	wg.Add(2)
	go producer()
	go producer()

	wg.Wait()

	close(ch)
	for val := range ch {
		slog.Info("successfully received", "value", val)
	}
}
