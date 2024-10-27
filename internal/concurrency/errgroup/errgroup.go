package errgroup

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func Pitfall() {
	g, ctx := errgroup.WithContext(context.Background())

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to goroutine workshop\n")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Starting Gin server
	g.Go(func() error {
		err := srv.ListenAndServe()
		slog.ErrorContext(ctx, "Gin ListenAndServe routine", "error", err)
		return err
	})

	// Another task
	g.Go(func() error {
		// Simulate some task that might return an error
		time.Sleep(5 * time.Second)
		return errors.New("simulated error in another task")
	})

	// Wait for goroutines to finish and handle errors
	if err := g.Wait(); err != nil {
		slog.ErrorContext(ctx, "errgroup.Wait()", "error", err)
	}
}

func Server() {
	g, ctx := errgroup.WithContext(context.Background())

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Starting Gin server
	g.Go(func() error {
		err := srv.ListenAndServe()
		slog.ErrorContext(ctx, "Gin ListenAndServe routine", "error", err)
		return err
	})

	// Shutdown routine
	g.Go(func() error {
		<-ctx.Done()
		slog.WarnContext(ctx, "Shutdown routine", "context.Err()", ctx.Err())
		return srv.Shutdown(ctx)
	})

	// Another pattern.go task
	g.Go(func() error {
		// Simulate some task that might return an error
		time.Sleep(5 * time.Second)
		return errors.New("simulated error in another task")
	})

	// Wait for goroutines to finish and handle errors
	if err := g.Wait(); err != nil {
		slog.ErrorContext(ctx, "errgroup.Wait()", "error", err)
	}
}
