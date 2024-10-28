package once

import (
	"sync"
	"testing"
)

// TestBasicSyncOnce demonstrates the intermediate usage of sync.Once.
func TestBasicSyncOnce(t *testing.T) {
	var once sync.Once
	var count int

	increment := func() { count++ }
	once.Do(increment)
	once.Do(increment)

	if count != 1 {
		t.Errorf("Expected 1, got %v", count)
	}
}

var onceA, onceB sync.Once

func initA(wg *sync.WaitGroup) {
	defer wg.Done()
	onceB.Do(func() {
		initB(wg)
	})
}

func initB(wg *sync.WaitGroup) {
	defer wg.Done()
	onceA.Do(func() {
		initA(wg)
	})
}

// TestBadUsageSyncOnce demonstrates a bad usage of sync.Once.
func TestBadUsageSyncOnce(t *testing.T) {
	t.Skip("Comment out to demonstrate incorrect usage")
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go initA(wg)
	go initB(wg)

	wg.Wait()
}
