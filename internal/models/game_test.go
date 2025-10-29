package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGame_IsCompleted(t *testing.T) {
	tests := []struct {
		name     string
		status   string
		expected bool
	}{
		{"completed game", "completed", true},
		{"in progress game", "in_progress", false},
		{"scheduled game", "scheduled", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &Game{Status: tt.status}
			assert.Equal(t, tt.expected, game.IsCompleted())
		})
	}
}

func TestGame_GetWinner(t *testing.T) {
	homeTeam := Team{ID: 1, Name: "Lakers"}
	awayTeam := Team{ID: 2, Name: "Celtics"}

	tests := []struct {
		name        string
		status      string
		homeScore   int
		awayScore   int
		expectedWin *Team
	}{
		{"home team wins", "completed", 110, 108, &homeTeam},
		{"away team wins", "completed", 108, 110, &awayTeam},
		{"tie game", "completed", 110, 110, nil},
		{"game not completed", "in_progress", 110, 108, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &Game{
				HomeTeam:  homeTeam,
				AwayTeam:  awayTeam,
				HomeScore: tt.homeScore,
				AwayScore: tt.awayScore,
				Status:    tt.status,
			}

			winner := game.GetWinner()
			if tt.expectedWin == nil {
				assert.Nil(t, winner)
			} else {
				assert.NotNil(t, winner)
				assert.Equal(t, tt.expectedWin.ID, winner.ID)
			}
		})
	}
}