package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeremielumandong/nba-result/internal/models"
	"github.com/jeremielumandong/nba-result/internal/services"
)

// GetTeams handles GET /api/v1/teams
func GetTeams(w http.ResponseWriter, r *http.Request) {
	teamService := services.NewTeamService()
	teams, err := teamService.GetAllTeams()
	if err != nil {
		http.Error(w, "Failed to retrieve teams", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(teams); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// GetTeamByID handles GET /api/v1/teams/{id}
func GetTeamByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	teamService := services.NewTeamService()
	team, err := teamService.GetTeamByID(id)
	if err != nil {
		if err == models.ErrTeamNotFound {
			http.Error(w, "Team not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve team", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(team); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}