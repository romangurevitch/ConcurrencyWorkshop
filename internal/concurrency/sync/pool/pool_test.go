package pool

import (
	"bytes"
	"sync"
	"testing"
)

// bufferPool is a pool of byte buffers.
var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

// TestBufferPool demonstrates a realistic example of using sync.Pool to pool byte buffers.
func TestBufferPool(t *testing.T) {
	// Acquire a buffer from the pool.
	buf, ok := bufferPool.Get().(*bytes.Buffer)
	if !ok {
		t.Errorf("Not a buffer type")
	}
	// Use the buffer.
	buf.WriteString("Hello, World!")
	if buf.String() != "Hello, World!" {
		t.Errorf("Expected 'Hello, World!', got '%s'", buf.String())
	}
	// Reset the buffer and put it back into the pool.
	buf.Reset()
	bufferPool.Put(buf)

	// Acquire the buffer again and ensure it's reset.
	buf, ok = bufferPool.Get().(*bytes.Buffer)
	if !ok {
		t.Errorf("Not a buffer type")
	}
	if buf.Len() != 0 {
		t.Errorf("Expected 0, got %v", buf.Len())
	}
}
