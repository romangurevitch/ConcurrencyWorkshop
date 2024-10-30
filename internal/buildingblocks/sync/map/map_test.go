package syncmap

import (
	"sync"
	"testing"
)

// TestBasicSyncMap demonstrates intermediate usage of sync.Map.
func TestBasicSyncMap(t *testing.T) {
	var m sync.Map

	m.Store("key1", "value1")
	m.Store("key2", "value2")

	value, ok := m.Load("key1")
	if !ok || value.(string) != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}

	value, ok = m.LoadOrStore("key2", "value3")
	if !ok || value.(string) != "value2" {
		t.Errorf("Expected value2, got %v", value)
	}

	value, ok = m.LoadOrStore("key3", "value3")
	if ok || value.(string) != "value3" {
		t.Errorf("Expected value3, got %v", value)
	}
}

// TestSyncMapRange demonstrates the usage of Range method of sync.Map.
func TestSyncMapRange(t *testing.T) {
	var m sync.Map

	m.Store("key1", "value1")
	m.Store("key2", "value2")

	var keys []string
	m.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})

	if len(keys) != 2 {
		t.Errorf("Expected 2 keys, got %v", keys)
	}
}
