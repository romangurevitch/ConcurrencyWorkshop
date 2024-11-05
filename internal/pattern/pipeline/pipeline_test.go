package pipeline

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ErrAtValue3 = errors.New("error at value 3")

func TestGeneratorAndProcessor(t *testing.T) {
	tests := []struct {
		name           string
		max            int
		processFunc    ProcessFunc[int, string]
		cancel         bool // whether to cancel the context before waiting for the result
		expectedOutput []Result[string]
	}{
		{
			name: "Successful operation",
			max:  5,
			processFunc: func(ctx context.Context, res Result[int]) Result[string] {
				return Result[string]{Value: "Processed: " + strconv.Itoa(res.Value)}
			},
			expectedOutput: []Result[string]{
				{Value: "Processed: 0"},
				{Value: "Processed: 1"},
				{Value: "Processed: 2"},
				{Value: "Processed: 3"},
				{Value: "Processed: 4"},
			},
		},
		{
			name: "Context cancellation",
			max:  5,
			processFunc: func(ctx context.Context, res Result[int]) Result[string] {
				return Result[string]{Value: "Processed: " + strconv.Itoa(res.Value)}
			},
			cancel:         true,
			expectedOutput: nil, // No output expected due to context cancellation.
		},
		{
			name: "Processing error",
			max:  5,
			processFunc: func(ctx context.Context, res Result[int]) Result[string] {
				if res.Value == 3 {
					return Result[string]{Err: ErrAtValue3}
				}
				return Result[string]{Value: "Processed: " + strconv.Itoa(res.Value)}
			},
			expectedOutput: []Result[string]{
				{Value: "Processed: 0"},
				{Value: "Processed: 1"},
				{Value: "Processed: 2"},
				{Err: ErrAtValue3},
				{Value: "Processed: 4"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			if tt.cancel {
				cancel() // cancel the context before waiting for the result
			}

			inputCh := make(chan Result[int])
			go func() {
				for i := 0; i < tt.max; i++ {
					inputCh <- Result[int]{Value: i}
				}
				close(inputCh)
			}()
			procCh := Pipe(ctx, inputCh, tt.processFunc)

			var gotResults []Result[string]
			for res := range procCh {
				gotResults = append(gotResults, res)
			}

			assert.ElementsMatch(t, tt.expectedOutput, gotResults)
		})
	}
}
