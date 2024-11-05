package test

import (
	"log/slog"
	"os"
	"testing"
	"time"
)

func ExitAfter(duration time.Duration) {
	go func() {
		<-time.After(duration)
		slog.Error("timeout exceeded, terminating program.")
		os.Exit(1)
	}()
}

func ExpectPanic(t *testing.T) {
	if err := recover(); err == nil {
		t.Fatal("Expected a panic!")
	}
}

func ExpectNoPanic(t *testing.T) {
	if err := recover(); err != nil {
		t.Fatal("Don’t panic and always carry a towel!", "error:", err)
	}
}
