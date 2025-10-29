// Package services contains business logic for NBA-related operations.
package services

import (
	"fmt"
	"time"

	"github.com/jeremielumandong/nba-result/internal/models"
	"github.com/jeremielumandong/nba-result/internal/repository"
)

// NBAServiceInterface defines the contract for NBA-related business operations.
type NBAServiceInterface interface {
	GetGames(dateFilter *time.Time, limit int) ([]*models.Game, error)
	GetGameByID(gameID string) (*models.Game, error)
	GetGamesByTeam(teamID string, limit int) ([]*models.Game, error)
	GetLiveGames() ([]*models.Game, error)
}

// NBAService implements NBA-related business logic.
type NBAService struct {
	repo repository.NBARepositoryInterface
}

// NewNBAService creates a new instance of NBAService.
func NewNBAService(repo repository.NBARepositoryInterface) NBAServiceInterface {
	return &NBAService{
		repo: repo,
	}
}

// GetGames retrieves games with optional date filtering and limit.
func (s *NBAService) GetGames(dateFilter *time.Time, limit int) ([]*models.Game, error) {
	if limit <= 0 {
		return nil, fmt.Errorf("limit must be greater than 0")
	}

	if limit > 100 {
		limit = 100 // Cap the limit to prevent excessive load
	}

	games, err := s.repo.GetAllGames(dateFilter, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve games: %w", err)
	}

	// Apply business logic - filter out canceled games unless specifically requested
	filteredGames := s.filterCanceledGames(games)

	return filteredGames, nil
}

// GetGameByID retrieves a specific game by its ID.
func (s *NBAService) GetGameByID(gameID string) (*models.Game, error) {
	if gameID == "" {
		return nil, fmt.Errorf("game ID cannot be empty")
	}

	game, err := s.repo.GetGameByID(gameID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve game: %w", err)
	}

	if game == nil {
		return nil, fmt.Errorf("game not found")
	}

	return game, nil
}

// GetGamesByTeam retrieves games for a specific team.
func (s *NBAService) GetGamesByTeam(teamID string, limit int) ([]*models.Game, error) {
	if teamID == "" {
		return nil, fmt.Errorf("team ID cannot be empty")
	}

	if limit <= 0 {
		limit = 20 // Default limit for team games
	}

	if limit > 50 {
		limit = 50 // Cap the limit for team games
	}

	games, err := s.repo.GetGamesByTeam(teamID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve games for team %s: %w", teamID, err)
	}

	// Sort games by date (most recent first)
	sortedGames := s.sortGamesByDate(games)

	return sortedGames, nil
}

// GetLiveGames retrieves all currently live games.
func (s *NBAService) GetLiveGames() ([]*models.Game, error) {
	allGames, err := s.repo.GetAllGames(nil, 100) // Get recent games to check for live ones
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve games: %w", err)
	}

	var liveGames []*models.Game
	for _, game := range allGames {
		if game.IsLive() {
			liveGames = append(liveGames, game)
		}
	}

	return liveGames, nil
}

// filterCanceledGames removes canceled games from the slice unless specifically needed.
func (s *NBAService) filterCanceledGames(games []*models.Game) []*models.Game {
	var filtered []*models.Game
	for _, game := range games {
		if game.Status != string(models.GameStatusCanceled) {
			filtered = append(filtered, game)
		}
	}
	return filtered
}

// sortGamesByDate sorts games by date in descending order (most recent first).
func (s *NBAService) sortGamesByDate(games []*models.Game) []*models.Game {
	// Simple bubble sort for demonstration - in production, use sort.Slice
	for i := 0; i < len(games)-1; i++ {
		for j := 0; j < len(games)-i-1; j++ {
			if games[j].GameDate.Before(games[j+1].GameDate) {
				games[j], games[j+1] = games[j+1], games[j]
			}
		}
	}
	return games
}

// validateGameData ensures game data meets business requirements.
func (s *NBAService) validateGameData(game *models.Game) error {
	if game == nil {
		return fmt.Errorf("game cannot be nil")
	}

	if game.ID == "" {
		return fmt.Errorf("game ID is required")
	}

	if game.HomeTeam.ID == "" || game.AwayTeam.ID == "" {
		return fmt.Errorf("both home and away team IDs are required")
	}

	if !models.IsValidStatus(game.Status) {
		return fmt.Errorf("invalid game status: %s", game.Status)
	}

	if game.GameDate.IsZero() {
		return fmt.Errorf("game date is required")
	}

	return nil
}