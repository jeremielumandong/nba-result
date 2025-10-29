package services

import (
	"testing"

	"github.com/jeremielumandong/nba-result/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGameService_GetAllGames(t *testing.T) {
	service := NewGameService()

	games, err := service.GetAllGames()

	assert.NoError(t, err)
	assert.NotEmpty(t, games)
	assert.Greater(t, len(games), 0)
}

func TestGameService_GetGameByID_Found(t *testing.T) {
	service := NewGameService()

	game, err := service.GetGameByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, game)
	assert.Equal(t, 1, game.ID)
}

func TestGameService_GetGameByID_NotFound(t *testing.T) {
	service := NewGameService()

	game, err := service.GetGameByID(999)

	assert.Error(t, err)
	assert.Equal(t, models.ErrGameNotFound, err)
	assert.Nil(t, game)
}

func TestGameService_GetGamesByTeam(t *testing.T) {
	service := NewGameService()

	games, err := service.GetGamesByTeam(1)

	assert.NoError(t, err)
	assert.NotEmpty(t, games)

	// Verify all games include the specified team
	for _, game := range games {
		assert.True(t, game.HomeTeam.ID == 1 || game.AwayTeam.ID == 1)
	}
}