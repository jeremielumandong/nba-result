// Package models contains data structures and domain models
package models

import (
	"errors"
	"time"
)

// Common errors
var (
	ErrGameNotFound = errors.New("game not found")
	ErrTeamNotFound = errors.New("team not found")
)

// Game represents an NBA game
type Game struct {
	ID          int       `json:"id"`
	HomeTeam    Team      `json:"home_team"`
	AwayTeam    Team      `json:"away_team"`
	HomeScore   int       `json:"home_score"`
	AwayScore   int       `json:"away_score"`
	GameDate    time.Time `json:"game_date"`
	Status      string    `json:"status"` // "scheduled", "in_progress", "completed"
	Season      string    `json:"season"`
	GameType    string    `json:"game_type"` // "regular", "playoff", "preseason"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// IsCompleted returns true if the game is completed
func (g *Game) IsCompleted() bool {
	return g.Status == "completed"
}

// GetWinner returns the winning team, or nil if game is not completed or tied
func (g *Game) GetWinner() *Team {
	if !g.IsCompleted() {
		return nil
	}

	if g.HomeScore > g.AwayScore {
		return &g.HomeTeam
	} else if g.AwayScore > g.HomeScore {
		return &g.AwayTeam
	}

	return nil // Tie game
}