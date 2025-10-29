package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGame_GetWinner(t *testing.T) {
	game := &Game{
		ID:     "test-game",
		Status: string(GameStatusFinished),
		HomeTeam: Team{
			ID:   "home",
			Name: "Home Team",
		},
		AwayTeam: Team{
			ID:   "away",
			Name: "Away Team",
		},
	}

	// Test home team wins
	t.Run("home team wins", func(t *testing.T) {
		game.HomeScore = 110
		game.AwayScore = 105

		winner := game.GetWinner()
		assert.NotNil(t, winner)
		assert.Equal(t, "home", winner.ID)
	})

	// Test away team wins
	t.Run("away team wins", func(t *testing.T) {
		game.HomeScore = 95
		game.AwayScore = 102

		winner := game.GetWinner()
		assert.NotNil(t, winner)
		assert.Equal(t, "away", winner.ID)
	})

	// Test tie game
	t.Run("tie game", func(t *testing.T) {
		game.HomeScore = 100
		game.AwayScore = 100

		winner := game.GetWinner()
		assert.Nil(t, winner)
	})

	// Test game not finished
	t.Run("game not finished", func(t *testing.T) {
		game.Status = string(GameStatusLive)
		game.HomeScore = 110
		game.AwayScore = 105

		winner := game.GetWinner()
		assert.Nil(t, winner)
	})
}

func TestGame_IsLive(t *testing.T) {
	game := &Game{
		ID: "test-game",
	}

	// Test live game
	t.Run("live game", func(t *testing.T) {
		game.Status = string(GameStatusLive)
		assert.True(t, game.IsLive())
	})

	// Test finished game
	t.Run("finished game", func(t *testing.T) {
		game.Status = string(GameStatusFinished)
		assert.False(t, game.IsLive())
	})

	// Test scheduled game
	t.Run("scheduled game", func(t *testing.T) {
		game.Status = string(GameStatusScheduled)
		assert.False(t, game.IsLive())
	})
}

func TestGame_IsFinished(t *testing.T) {
	game := &Game{
		ID: "test-game",
	}

	// Test finished game
	t.Run("finished game", func(t *testing.T) {
		game.Status = string(GameStatusFinished)
		assert.True(t, game.IsFinished())
	})

	// Test live game
	t.Run("live game", func(t *testing.T) {
		game.Status = string(GameStatusLive)
		assert.False(t, game.IsFinished())
	})

	// Test scheduled game
	t.Run("scheduled game", func(t *testing.T) {
		game.Status = string(GameStatusScheduled)
		assert.False(t, game.IsFinished())
	})
}

func TestIsValidStatus(t *testing.T) {
	testCases := []struct {
		name     string
		status   string
		expected bool
	}{
		{
			name:     "valid scheduled status",
			status:   string(GameStatusScheduled),
			expected: true,
		},
		{
			name:     "valid live status",
			status:   string(GameStatusLive),
			expected: true,
		},
		{
			name:     "valid finished status",
			status:   string(GameStatusFinished),
			expected: true,
		},
		{
			name:     "valid postponed status",
			status:   string(GameStatusPostponed),
			expected: true,
		},
		{
			name:     "valid canceled status",
			status:   string(GameStatusCanceled),
			expected: true,
		},
		{
			name:     "invalid status",
			status:   "INVALID_STATUS",
			expected: false,
		},
		{
			name:     "empty status",
			status:   "",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsValidStatus(tc.status)
			assert.Equal(t, tc.expected, result)
		})
	}
}