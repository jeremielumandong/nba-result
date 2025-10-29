// Package api handles HTTP routing and handlers
package api

import (
	"github.com/gorilla/mux"
	"github.com/jeremielumandong/nba-result/internal/handlers"
)

// NewRouter creates and configures the main application router
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// API v1 routes
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	api.HandleFunc("/games", handlers.GetGames).Methods("GET")
	api.HandleFunc("/games/{id}", handlers.GetGameByID).Methods("GET")
	api.HandleFunc("/teams", handlers.GetTeams).Methods("GET")
	api.HandleFunc("/teams/{id}", handlers.GetTeamByID).Methods("GET")

	return r
}