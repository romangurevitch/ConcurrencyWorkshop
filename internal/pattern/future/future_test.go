package future

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var ErrTest = errors.New("test error")

func TestFuture(t *testing.T) {
	type args struct {
		computation ProcessFunc[int]
	}
	type testCase struct {
		name   string
		args   args
		cancel bool        // whether to cancel the context before waiting for the result
		want   Result[int] // expected result
	}

	tests := []testCase{
		{
			name: "Successful Computation",
			args: args{
				computation: func(ctx context.Context) (int, error) {
					return 42, nil
				},
			},
			want: Result[int]{Value: 42},
		},
		{
			name: "Computation Error",
			args: args{
				computation: func(ctx context.Context) (int, error) {
					return 0, ErrTest
				},
			},
			want: Result[int]{Err: ErrTest},
		},
		{
			name: "Context Cancellation",
			args: args{
				computation: func(ctx context.Context) (int, error) {
					time.Sleep(200 * time.Millisecond) // simulate a longer computation
					return 42, nil
				},
			},
			cancel: true,
			want:   Result[int]{Err: context.Canceled},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel() // ensure resources are cleaned up

			if tt.cancel {
				cancel() // cancel the context before waiting for the result
			}

			future := NewFuture(ctx, tt.args.computation)
			got := future.Result()

			assert.Equal(t, tt.want, got)
		})
	}
}
