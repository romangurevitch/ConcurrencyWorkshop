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

func CheckNoPanic(t *testing.T) {
	if err := recover(); err != nil {
		t.Fatal("Don’t panic and always carry a towel!", "error:", err)
	}
}
