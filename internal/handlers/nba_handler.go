// Package handlers contains HTTP request handlers for the NBA Results API.
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jeremielumandong/nba-result/internal/services"
	"github.com/jeremielumandong/nba-result/pkg/utils"
)

// NBAHandler handles HTTP requests related to NBA games.
type NBAHandler struct {
	nbaService services.NBAServiceInterface
}

// NewNBAHandler creates a new instance of NBAHandler.
func NewNBAHandler(nbaService services.NBAServiceInterface) *NBAHandler {
	return &NBAHandler{
		nbaService: nbaService,
	}
}

// SetupRoutes configures and returns the HTTP router with all NBA-related routes.
func (h *NBAHandler) SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Game routes
	api.HandleFunc("/games", h.GetAllGames).Methods("GET")
	api.HandleFunc("/games/{id}", h.GetGameByID).Methods("GET")
	api.HandleFunc("/teams/{team}/games", h.GetGamesByTeam).Methods("GET")

	// Health check
	api.HandleFunc("/health", h.HealthCheck).Methods("GET")

	// Add middleware
	api.Use(h.loggingMiddleware)
	api.Use(h.corsMiddleware)

	return router
}

// GetAllGames handles GET /api/games - retrieves all games with optional filtering.
func (h *NBAHandler) GetAllGames(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	// Parse optional date filter
	var dateFilter *time.Time
	if dateStr := queryParams.Get("date"); dateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", dateStr); err == nil {
			dateFilter = &parsedDate
		} else {
			h.sendErrorResponse(w, http.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD")
			return
		}
	}

	// Parse optional limit
	limit := 50 // default limit
	if limitStr := queryParams.Get("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 && parsedLimit <= 100 {
			limit = parsedLimit
		} else {
			h.sendErrorResponse(w, http.StatusBadRequest, "Invalid limit. Must be between 1 and 100")
			return
		}
	}

	games, err := h.nbaService.GetGames(dateFilter, limit)
	if err != nil {
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve games")
		return
	}

	h.sendJSONResponse(w, http.StatusOK, games)
}

// GetGameByID handles GET /api/games/{id} - retrieves a specific game by ID.
func (h *NBAHandler) GetGameByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["id"]

	if gameID == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Game ID is required")
		return
	}

	game, err := h.nbaService.GetGameByID(gameID)
	if err != nil {
		if err.Error() == "game not found" {
			h.sendErrorResponse(w, http.StatusNotFound, "Game not found")
		} else {
			h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve game")
		}
		return
	}

	h.sendJSONResponse(w, http.StatusOK, game)
}

// GetGamesByTeam handles GET /api/teams/{team}/games - retrieves games for a specific team.
func (h *NBAHandler) GetGamesByTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["team"]

	if teamID == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Team ID is required")
		return
	}

	queryParams := r.URL.Query()
	limit := 20 // default limit for team games
	if limitStr := queryParams.Get("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 && parsedLimit <= 50 {
			limit = parsedLimit
		}
	}

	games, err := h.nbaService.GetGamesByTeam(teamID, limit)
	if err != nil {
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve team games")
		return
	}

	h.sendJSONResponse(w, http.StatusOK, games)
}

// HealthCheck handles GET /api/health - returns the health status of the API.
func (h *NBAHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	healthStatus := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"version":   "1.0.0",
	}

	h.sendJSONResponse(w, http.StatusOK, healthStatus)
}

// sendJSONResponse sends a JSON response with the given status code and data.
func (h *NBAHandler) sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// sendErrorResponse sends an error response with the given status code and message.
func (h *NBAHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := map[string]interface{}{
		"error":     message,
		"timestamp": time.Now().UTC(),
		"status":    statusCode,
	}

	h.sendJSONResponse(w, statusCode, errorResponse)
}

// loggingMiddleware logs HTTP requests.
func (h *NBAHandler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		utils.LogInfo(fmt.Sprintf("%s %s - %v", r.Method, r.RequestURI, time.Since(start)))
	})
}

// corsMiddleware adds CORS headers.
func (h *NBAHandler) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}