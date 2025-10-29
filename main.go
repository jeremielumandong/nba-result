package main

import (
	"log"
	"net/http"

	"github.com/jeremielumandong/nba-result/internal/api"
	"github.com/jeremielumandong/nba-result/internal/config"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize API router
	router := api.NewRouter()

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}