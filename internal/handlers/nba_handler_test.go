package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/jeremielumandong/nba-result/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockNBAService is a mock implementation of NBAServiceInterface for testing.
type MockNBAService struct {
	mock.Mock
}

func (m *MockNBAService) GetGames(dateFilter *time.Time, limit int) ([]*models.Game, error) {
	args := m.Called(dateFilter, limit)
	return args.Get(0).([]*models.Game), args.Error(1)
}

func (m *MockNBAService) GetGameByID(gameID string) (*models.Game, error) {
	args := m.Called(gameID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Game), args.Error(1)
}

func (m *MockNBAService) GetGamesByTeam(teamID string, limit int) ([]*models.Game, error) {
	args := m.Called(teamID, limit)
	return args.Get(0).([]*models.Game), args.Error(1)
}

func (m *MockNBAService) GetLiveGames() ([]*models.Game, error) {
	args := m.Called()
	return args.Get(0).([]*models.Game), args.Error(1)
}

func TestNBAHandler_GetAllGames(t *testing.T) {
	mockService := new(MockNBAService)
	handler := NewNBAHandler(mockService)

	// Test successful request
	t.Run("successful request", func(t *testing.T) {
		mockGames := []*models.Game{
			{
				ID:     "game1",
				Status: string(models.GameStatusFinished),
			},
		}

		mockService.On("GetGames", (*time.Time)(nil), 50).Return(mockGames, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/games", nil)
		w := httptest.NewRecorder()

		handler.GetAllGames(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []*models.Game
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 1)
		assert.Equal(t, "game1", response[0].ID)

		mockService.AssertExpectations(t)
	})

	// Test with date filter
	t.Run("with date filter", func(t *testing.T) {
		mockService = new(MockNBAService) // Reset mock
		handler = NewNBAHandler(mockService)

		testDate := time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC)
		mockGames := []*models.Game{}

		mockService.On("GetGames", &testDate, 50).Return(mockGames, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/games?date=2023-12-15", nil)
		w := httptest.NewRecorder()

		handler.GetAllGames(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})

	// Test invalid date format
	t.Run("invalid date format", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/games?date=invalid-date", nil)
		w := httptest.NewRecorder()

		handler.GetAllGames(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestNBAHandler_GetGameByID(t *testing.T) {
	mockService := new(MockNBAService)
	handler := NewNBAHandler(mockService)

	// Test successful request
	t.Run("successful request", func(t *testing.T) {
		mockGame := &models.Game{
			ID:     "game1",
			Status: string(models.GameStatusFinished),
		}

		mockService.On("GetGameByID", "game1").Return(mockGame, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/games/game1", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "game1"})
		w := httptest.NewRecorder()

		handler.GetGameByID(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Game
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "game1", response.ID)

		mockService.AssertExpectations(t)
	})

	// Test game not found
	t.Run("game not found", func(t *testing.T) {
		mockService = new(MockNBAService)
		handler = NewNBAHandler(mockService)

		mockService.On("GetGameByID", "nonexistent").Return((*models.Game)(nil), assert.AnError)

		req := httptest.NewRequest(http.MethodGet, "/api/games/nonexistent", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "nonexistent"})
		w := httptest.NewRecorder()

		handler.GetGameByID(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockService.AssertExpectations(t)
	})
}

func TestNBAHandler_GetGamesByTeam(t *testing.T) {
	mockService := new(MockNBAService)
	handler := NewNBAHandler(mockService)

	// Test successful request
	t.Run("successful request", func(t *testing.T) {
		mockGames := []*models.Game{
			{
				ID:     "game1",
				Status: string(models.GameStatusFinished),
			},
		}

		mockService.On("GetGamesByTeam", "lal", 20).Return(mockGames, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/teams/lal/games", nil)
		req = mux.SetURLVars(req, map[string]string{"team": "lal"})
		w := httptest.NewRecorder()

		handler.GetGamesByTeam(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []*models.Game
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 1)

		mockService.AssertExpectations(t)
	})
}

func TestNBAHandler_HealthCheck(t *testing.T) {
	mockService := new(MockNBAService)
	handler := NewNBAHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	w := httptest.NewRecorder()

	handler.HealthCheck(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "healthy", response["status"])
	assert.Equal(t, "1.0.0", response["version"])
}