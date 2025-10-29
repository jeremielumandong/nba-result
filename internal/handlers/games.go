package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeremielumandong/nba-result/internal/models"
	"github.com/jeremielumandong/nba-result/internal/services"
)

// GetGames handles GET /api/v1/games
func GetGames(w http.ResponseWriter, r *http.Request) {
	gameService := services.NewGameService()
	games, err := gameService.GetAllGames()
	if err != nil {
		http.Error(w, "Failed to retrieve games", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(games); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// GetGameByID handles GET /api/v1/games/{id}
func GetGameByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid game ID", http.StatusBadRequest)
		return
	}

	gameService := services.NewGameService()
	game, err := gameService.GetGameByID(id)
	if err != nil {
		if err == models.ErrGameNotFound {
			http.Error(w, "Game not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve game", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(game); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}