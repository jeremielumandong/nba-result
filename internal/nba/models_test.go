package nba

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame_Winner(t *testing.T) {
	tests := []struct {
		name     string
		game     Game
		expected string
	}{
		{
			name: "Home team wins",
			game: Game{
				HomeTeam: Team{Name: "Lakers", Score: 110},
				AwayTeam: Team{Name: "Warriors", Score: 105},
			},
			expected: "Lakers",
		},
		{
			name: "Away team wins",
			game: Game{
				HomeTeam: Team{Name: "Lakers", Score: 105},
				AwayTeam: Team{Name: "Warriors", Score: 110},
			},
			expected: "Warriors",
		},
		{
			name: "Tie game",
			game: Game{
				HomeTeam: Team{Name: "Lakers", Score: 105},
				AwayTeam: Team{Name: "Warriors", Score: 105},
			},
			expected: "TIE",
		},
		{
			name: "Zero scores",
			game: Game{
				HomeTeam: Team{Name: "Lakers", Score: 0},
				AwayTeam: Team{Name: "Warriors", Score: 0},
			},
			expected: "TIE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.game.Winner()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGame_IsFinished(t *testing.T) {
	tests := []struct {
		name     string
		status   string
		expected bool
	}{
		{"Final status", "Final", true},
		{"3 - Final status", "3 - Final", true},
		{"In progress", "2nd Qtr", false},
		{"Halftime", "Halftime", false},
		{"Not started", "7:30 pm ET", false},
		{"Empty status", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := Game{Status: tt.status}
			result := game.IsFinished()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGame_ToSummary(t *testing.T) {
	game := Game{
		GameDate: "2024-01-15",
		HomeTeam: Team{Name: "Lakers", Score: 110},
		AwayTeam: Team{Name: "Warriors", Score: 105},
		Status:   "Final",
	}

	summary := game.ToSummary()

	assert.Equal(t, "Warriors vs Lakers", summary.Matchup)
	assert.Equal(t, "105 - 110", summary.FinalScore)
	assert.Equal(t, "Lakers", summary.Winner)
	assert.Equal(t, "Final", summary.Status)
	assert.Equal(t, "2024-01-15", summary.Date)
}