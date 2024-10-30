package once

import (
	"sync"
	"testing"
)

func BenchmarkOnce(b *testing.B) {
	once := sync.Once{}
	once.Do(func() {
		// Call the function once
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		once.Do(func() {
			// Skipping
		})
	}
}
