package pipeline

import (
	"context"
	"log/slog"
)

// Result is a generic type to encapsulate the result of an operation.
type Result[T any] struct {
	Value T
	Err   error
}

// ProcessFunc defines a function type that processes a Result of type T and produces a Result of type U.
type ProcessFunc[T any, U any] func(context.Context, Result[T]) Result[U]

// Pipe reads Results of type T from inCh, processes them using the provided operation op,
// and sends the Results of type U on a new channel.
func Pipe[T any, U any](ctx context.Context, inCh <-chan Result[T], processFunc ProcessFunc[T, U]) <-chan Result[U] {
	outCh := make(chan Result[U])
	go func() {
		defer close(outCh) // Ensure the channel is closed when the goroutine exits.
		for {
			select {
			case <-ctx.Done():
				slog.Info("shutting down goroutine", "reason", ctx.Err())
				return
			case in, ok := <-inCh:
				if !ok {
					return // jobs channel closed, exit worker
				}
				outCh <- processFunc(ctx, in) // Process the result using processFunc and send it on the output channel.
			}
		}
	}()
	return outCh
}
