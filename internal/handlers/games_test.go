package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jeremielumandong/nba-result/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetGames(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/games", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	GetGames(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var games []models.Game
	err = json.Unmarshal(rr.Body.Bytes(), &games)
	assert.NoError(t, err)
	assert.NotEmpty(t, games)
}

func TestGetGameByID_ValidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/games/1", nil)
	assert.NoError(t, err)

	// Set up mux vars
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	rr := httptest.NewRecorder()

	GetGameByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var game models.Game
	err = json.Unmarshal(rr.Body.Bytes(), &game)
	assert.NoError(t, err)
	assert.Equal(t, 1, game.ID)
}

func TestGetGameByID_InvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/games/invalid", nil)
	assert.NoError(t, err)

	req = mux.SetURLVars(req, map[string]string{"id": "invalid"})

	rr := httptest.NewRecorder()

	GetGameByID(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetGameByID_NotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/games/999", nil)
	assert.NoError(t, err)

	req = mux.SetURLVars(req, map[string]string{"id": "999"})

	rr := httptest.NewRecorder()

	GetGameByID(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}