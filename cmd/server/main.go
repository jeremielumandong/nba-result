// Package main is the entry point for the NBA Results API server.
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jeremielumandong/nba-result/internal/config"
	"github.com/jeremielumandong/nba-result/internal/handlers"
	"github.com/jeremielumandong/nba-result/internal/repository"
	"github.com/jeremielumandong/nba-result/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize dependencies
	repo := repository.NewNBARepository()
	svc := services.NewNBAService(repo)
	handler := handlers.NewNBAHandler(svc)

	// Setup HTTP server
	server := &http.Server{
		Addr:         cfg.ServerAddress,
		Handler:      handler.SetupRoutes(),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting NBA Results API server on %s", cfg.ServerAddress)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}