// Package services contains business logic and service layer
package services

import (
	"time"

	"github.com/jeremielumandong/nba-result/internal/models"
)

// GameService handles game-related business logic
type GameService struct {
	// In a real implementation, this would have a repository/database dependency
}

// NewGameService creates a new GameService instance
func NewGameService() *GameService {
	return &GameService{}
}

// GetAllGames returns all games
// TODO: Implement database integration
func (s *GameService) GetAllGames() ([]models.Game, error) {
	// Mock data for now - replace with actual database calls
	return []models.Game{
		{
			ID: 1,
			HomeTeam: models.Team{
				ID:           1,
				Name:         "Lakers",
				City:         "Los Angeles",
				Abbreviation: "LAL",
				Conference:   "Western",
				Division:     "Pacific",
			},
			AwayTeam: models.Team{
				ID:           2,
				Name:         "Celtics",
				City:         "Boston",
				Abbreviation: "BOS",
				Conference:   "Eastern",
				Division:     "Atlantic",
			},
			HomeScore: 110,
			AwayScore: 108,
			GameDate:  time.Now().AddDate(0, 0, -1),
			Status:    "completed",
			Season:    "2023-24",
			GameType:  "regular",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}

// GetGameByID returns a game by its ID
func (s *GameService) GetGameByID(id int) (*models.Game, error) {
	// Mock implementation - replace with actual database lookup
	games, err := s.GetAllGames()
	if err != nil {
		return nil, err
	}

	for _, game := range games {
		if game.ID == id {
			return &game, nil
		}
	}

	return nil, models.ErrGameNotFound
}

// GetGamesByTeam returns all games for a specific team
func (s *GameService) GetGamesByTeam(teamID int) ([]models.Game, error) {
	games, err := s.GetAllGames()
	if err != nil {
		return nil, err
	}

	var teamGames []models.Game
	for _, game := range games {
		if game.HomeTeam.ID == teamID || game.AwayTeam.ID == teamID {
			teamGames = append(teamGames, game)
		}
	}

	return teamGames, nil
}