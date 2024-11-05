package pubsub

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPubSub(t *testing.T) {
	tests := []struct {
		name   string
		action func(*PubSub[string], *testing.T)
	}{
		{
			name: "Single Subscribe and Publish",
			action: func(ps *PubSub[string], t *testing.T) {
				ch := make(chan Result[string], 1)
				ps.Subscribe("topic1", ch)
				ps.Publish("topic1", "message1")
				result := <-ch
				assert.Equal(t, "message1", result.Value, "they should be equal")
			},
		},
		{
			name: "Multiple Subscribe and Publish",
			action: func(ps *PubSub[string], t *testing.T) {
				ch1 := make(chan Result[string], 1)
				ch2 := make(chan Result[string], 1)
				ps.Subscribe("topic2", ch1)
				ps.Subscribe("topic2", ch2)
				ps.Publish("topic2", "message2")
				result1 := <-ch1
				result2 := <-ch2
				assert.Equal(t, "message2", result1.Value, "they should be equal")
				assert.Equal(t, "message2", result2.Value, "they should be equal")
			},
		},
		{
			name: "Unsubscribe",
			action: func(ps *PubSub[string], t *testing.T) {
				ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
				defer cancelFunc()

				ch := make(chan Result[string], 1)
				ps.Subscribe("topic3", ch)
				ps.Unsubscribe("topic3", ch)
				ps.Publish("topic3", "message3")
				select {
				case _, ok := <-ch:
					assert.False(t, ok, "expected channel to be closed or empty, but received a message")
				case <-ctx.Done():
				}
			},
		},
		{
			name: "Multiple Subscribe and Publish, non buffered channel",
			action: func(ps *PubSub[string], t *testing.T) {
				ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
				defer cancelFunc()

				ch1 := make(chan Result[string], 1)
				ch2 := make(chan Result[string])
				ps.Subscribe("topic2", ch1)
				ps.Subscribe("topic2", ch2)
				ps.Publish("topic2", "message2")
				result1 := <-ch1
				assert.Equal(t, "message2", result1.Value, "they should be equal")

				select {
				case _, ok := <-ch2:
					assert.False(t, ok, "expected channel to be closed or empty, but received a message")
				case <-ctx.Done():
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := NewPubSub[string]()
			tt.action(ps, t)
		})
	}
}
