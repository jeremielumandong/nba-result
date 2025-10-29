// Package models contains data structures used throughout the application.
package models

import (
	"time"
)

// Game represents an NBA game with all relevant information.
type Game struct {
	ID          string    `json:"id"`
	HomeTeam    Team      `json:"home_team"`
	AwayTeam    Team      `json:"away_team"`
	GameDate    time.Time `json:"game_date"`
	Status      string    `json:"status"`
	HomeScore   int       `json:"home_score"`
	AwayScore   int       `json:"away_score"`
	Period      int       `json:"period"`
	GameClock   string    `json:"game_clock,omitempty"`
	Arena       Arena     `json:"arena"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Team represents an NBA team.
type Team struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	City         string `json:"city"`
	Abbreviation string `json:"abbreviation"`
	Logo         string `json:"logo,omitempty"`
}

// Arena represents the venue where the game is played.
type Arena struct {
	Name     string `json:"name"`
	City     string `json:"city"`
	State    string `json:"state"`
	Capacity int    `json:"capacity,omitempty"`
}

// GameStatus defines possible game statuses.
type GameStatus string

const (
	GameStatusScheduled GameStatus = "SCHEDULED"
	GameStatusLive      GameStatus = "LIVE"
	GameStatusFinished  GameStatus = "FINISHED"
	GameStatusPostponed GameStatus = "POSTPONED"
	GameStatusCanceled  GameStatus = "CANCELED"
)

// IsValidStatus checks if the provided status is valid.
func IsValidStatus(status string) bool {
	validStatuses := []GameStatus{
		GameStatusScheduled,
		GameStatusLive,
		GameStatusFinished,
		GameStatusPostponed,
		GameStatusCanceled,
	}

	for _, validStatus := range validStatuses {
		if string(validStatus) == status {
			return true
		}
	}
	return false
}

// GetWinner returns the winning team or nil if the game is not finished or is a tie.
func (g *Game) GetWinner() *Team {
	if g.Status != string(GameStatusFinished) {
		return nil
	}

	if g.HomeScore > g.AwayScore {
		return &g.HomeTeam
	} else if g.AwayScore > g.HomeScore {
		return &g.AwayTeam
	}

	return nil // Tie game
}

// IsLive returns true if the game is currently being played.
func (g *Game) IsLive() bool {
	return g.Status == string(GameStatusLive)
}

// IsFinished returns true if the game has ended.
func (g *Game) IsFinished() bool {
	return g.Status == string(GameStatusFinished)
}