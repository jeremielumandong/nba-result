package services

import (
	"errors"
	"testing"
	"time"

	"github.com/jeremielumandong/nba-result/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockNBARepository is a mock implementation of NBARepositoryInterface for testing.
type MockNBARepository struct {
	mock.Mock
}

func (m *MockNBARepository) GetAllGames(dateFilter *time.Time, limit int) ([]*models.Game, error) {
	args := m.Called(dateFilter, limit)
	return args.Get(0).([]*models.Game), args.Error(1)
}

func (m *MockNBARepository) GetGameByID(gameID string) (*models.Game, error) {
	args := m.Called(gameID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Game), args.Error(1)
}

func (m *MockNBARepository) GetGamesByTeam(teamID string, limit int) ([]*models.Game, error) {
	args := m.Called(teamID, limit)
	return args.Get(0).([]*models.Game), args.Error(1)
}

func (m *MockNBARepository) SaveGame(game *models.Game) error {
	args := m.Called(game)
	return args.Error(0)
}

func (m *MockNBARepository) UpdateGame(game *models.Game) error {
	args := m.Called(game)
	return args.Error(0)
}

func (m *MockNBARepository) DeleteGame(gameID string) error {
	args := m.Called(gameID)
	return args.Error(0)
}

func TestNBAService_GetGames(t *testing.T) {
	mockRepo := new(MockNBARepository)
	service := NewNBAService(mockRepo)

	// Test successful retrieval
	t.Run("successful retrieval", func(t *testing.T) {
		mockGames := []*models.Game{
			{
				ID:     "game1",
				Status: string(models.GameStatusFinished),
			},
			{
				ID:     "game2",
				Status: string(models.GameStatusCanceled),
			},
		}

		mockRepo.On("GetAllGames", (*time.Time)(nil), 10).Return(mockGames, nil)

		games, err := service.GetGames(nil, 10)

		assert.NoError(t, err)
		assert.Len(t, games, 1) // Canceled game should be filtered out
		assert.Equal(t, "game1", games[0].ID)

		mockRepo.AssertExpectations(t)
	})

	// Test invalid limit
	t.Run("invalid limit", func(t *testing.T) {
		_, err := service.GetGames(nil, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "limit must be greater than 0")
	})

	// Test repository error
	t.Run("repository error", func(t *testing.T) {
		mockRepo = new(MockNBARepository)
		service = NewNBAService(mockRepo)

		mockRepo.On("GetAllGames", (*time.Time)(nil), 10).Return([]*models.Game(nil), errors.New("database error"))

		_, err := service.GetGames(nil, 10)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to retrieve games")
	})
}

func TestNBAService_GetGameByID(t *testing.T) {
	mockRepo := new(MockNBARepository)
	service := NewNBAService(mockRepo)

	// Test successful retrieval
	t.Run("successful retrieval", func(t *testing.T) {
		mockGame := &models.Game{
			ID:     "game1",
			Status: string(models.GameStatusFinished),
		}

		mockRepo.On("GetGameByID", "game1").Return(mockGame, nil)

		game, err := service.GetGameByID("game1")

		assert.NoError(t, err)
		assert.NotNil(t, game)
		assert.Equal(t, "game1", game.ID)

		mockRepo.AssertExpectations(t)
	})

	// Test empty game ID
	t.Run("empty game ID", func(t *testing.T) {
		_, err := service.GetGameByID("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "game ID cannot be empty")
	})

	// Test game not found
	t.Run("game not found", func(t *testing.T) {
		mockRepo = new(MockNBARepository)
		service = NewNBAService(mockRepo)

		mockRepo.On("GetGameByID", "nonexistent").Return((*models.Game)(nil), nil)

		_, err := service.GetGameByID("nonexistent")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "game not found")
	})
}

func TestNBAService_GetGamesByTeam(t *testing.T) {
	mockRepo := new(MockNBARepository)
	service := NewNBAService(mockRepo)

	// Test successful retrieval
	t.Run("successful retrieval", func(t *testing.T) {
		mockGames := []*models.Game{
			{
				ID:       "game1",
				Status:   string(models.GameStatusFinished),
				GameDate: time.Now().AddDate(0, 0, -1),
			},
			{
				ID:       "game2",
				Status:   string(models.GameStatusFinished),
				GameDate: time.Now(),
			},
		}

		mockRepo.On("GetGamesByTeam", "lal", 20).Return(mockGames, nil)

		games, err := service.GetGamesByTeam("lal", 20)

		assert.NoError(t, err)
		assert.Len(t, games, 2)
		// Games should be sorted by date (most recent first)
		assert.True(t, games[0].GameDate.After(games[1].GameDate) || games[0].GameDate.Equal(games[1].GameDate))

		mockRepo.AssertExpectations(t)
	})

	// Test empty team ID
	t.Run("empty team ID", func(t *testing.T) {
		_, err := service.GetGamesByTeam("", 20)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "team ID cannot be empty")
	})
}

func TestNBAService_GetLiveGames(t *testing.T) {
	mockRepo := new(MockNBARepository)
	service := NewNBAService(mockRepo)

	// Test successful retrieval with live games
	t.Run("successful retrieval with live games", func(t *testing.T) {
		mockGames := []*models.Game{
			{
				ID:     "game1",
				Status: string(models.GameStatusLive),
			},
			{
				ID:     "game2",
				Status: string(models.GameStatusFinished),
			},
			{
				ID:     "game3",
				Status: string(models.GameStatusLive),
			},
		}

		mockRepo.On("GetAllGames", (*time.Time)(nil), 100).Return(mockGames, nil)

		liveGames, err := service.GetLiveGames()

		assert.NoError(t, err)
		assert.Len(t, liveGames, 2) // Only live games should be returned
		for _, game := range liveGames {
			assert.True(t, game.IsLive())
		}

		mockRepo.AssertExpectations(t)
	})
}