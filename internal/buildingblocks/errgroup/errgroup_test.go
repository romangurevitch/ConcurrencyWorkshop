package errgroup

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

// TestNoError demonstrates a scenario where all tasks succeed.
func TestNoError(t *testing.T) {
	g, _ := errgroup.WithContext(context.Background())

	// Two tasks that succeed
	g.Go(func() error {
		return nil
	})
	g.Go(func() error {
		return nil
	})

	// Expecting no error from the group
	if err := g.Wait(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// TestBasic demonstrates intermediate usage of errgroup.
func TestBasicError(t *testing.T) {
	g, _ := errgroup.WithContext(context.Background())

	// Task that succeeds
	g.Go(func() error {
		return nil
	})

	errTask := errors.New("task error")
	// Task that fails
	g.Go(func() error {
		return errTask
	})

	// Expecting an error from the group
	if err := g.Wait(); err == nil {
		assert.ErrorIs(t, err, errTask)
	}
}

// TestDirectDefinition demonstrates a bad practice of defining errgroup.Group directly.
func TestDirectDefinition(t *testing.T) {
	// Not recommended: Defining errgroup.Group directly without a context
	var g errgroup.Group

	errTask := errors.New("task error")
	g.Go(func() error {
		return errTask
	})

	// This will work, but lacks context control which could lead to goroutine leaks or lack of context propagation.
	if err := g.Wait(); err == nil {
		assert.ErrorIs(t, err, errTask)
	}
}

func TestPitfall(t *testing.T) {
	t.Skip("Comment out to demonstrate error group incorrect usage")
	Pitfall()
}

func TestServer(t *testing.T) {
	Server()
}
