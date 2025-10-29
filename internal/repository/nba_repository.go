// Package repository handles data persistence and retrieval for NBA-related entities.
package repository

import (
	"fmt"
	"time"

	"github.com/jeremielumandong/nba-result/internal/models"
)

// NBARepositoryInterface defines the contract for NBA data operations.
type NBARepositoryInterface interface {
	GetAllGames(dateFilter *time.Time, limit int) ([]*models.Game, error)
	GetGameByID(gameID string) (*models.Game, error)
	GetGamesByTeam(teamID string, limit int) ([]*models.Game, error)
	SaveGame(game *models.Game) error
	UpdateGame(game *models.Game) error
	DeleteGame(gameID string) error
}

// NBARepository implements NBA data operations.
// In a real application, this would connect to a database or external API.
type NBARepository struct {
	// In-memory storage for demonstration purposes
	games map[string]*models.Game
}

// NewNBARepository creates a new instance of NBARepository.
func NewNBARepository() NBARepositoryInterface {
	repo := &NBARepository{
		games: make(map[string]*models.Game),
	}

	// Initialize with some sample data
	repo.initializeSampleData()

	return repo
}

// GetAllGames retrieves all games with optional date filtering and limit.
func (r *NBARepository) GetAllGames(dateFilter *time.Time, limit int) ([]*models.Game, error) {
	var games []*models.Game
	count := 0

	for _, game := range r.games {
		// Apply date filter if provided
		if dateFilter != nil {
			gameDate := game.GameDate.Truncate(24 * time.Hour)
			filterDate := dateFilter.Truncate(24 * time.Hour)
			if !gameDate.Equal(filterDate) {
				continue
			}
		}

		games = append(games, game)
		count++

		// Apply limit
		if limit > 0 && count >= limit {
			break
		}
	}

	return games, nil
}

// GetGameByID retrieves a specific game by its ID.
func (r *NBARepository) GetGameByID(gameID string) (*models.Game, error) {
	game, exists := r.games[gameID]
	if !exists {
		return nil, fmt.Errorf("game not found")
	}

	return game, nil
}

// GetGamesByTeam retrieves games for a specific team.
func (r *NBARepository) GetGamesByTeam(teamID string, limit int) ([]*models.Game, error) {
	var games []*models.Game
	count := 0

	for _, game := range r.games {
		// Check if the team is either home or away
		if game.HomeTeam.ID == teamID || game.AwayTeam.ID == teamID {
			games = append(games, game)
			count++

			// Apply limit
			if limit > 0 && count >= limit {
				break
			}
		}
	}

	return games, nil
}

// SaveGame saves a new game to the repository.
func (r *NBARepository) SaveGame(game *models.Game) error {
	if game == nil {
		return fmt.Errorf("game cannot be nil")
	}

	if game.ID == "" {
		return fmt.Errorf("game ID is required")
	}

	// Check if game already exists
	if _, exists := r.games[game.ID]; exists {
		return fmt.Errorf("game with ID %s already exists", game.ID)
	}

	// Set timestamps
	now := time.Now().UTC()
	game.CreatedAt = now
	game.UpdatedAt = now

	r.games[game.ID] = game
	return nil
}

// UpdateGame updates an existing game in the repository.
func (r *NBARepository) UpdateGame(game *models.Game) error {
	if game == nil {
		return fmt.Errorf("game cannot be nil")
	}

	if game.ID == "" {
		return fmt.Errorf("game ID is required")
	}

	// Check if game exists
	existingGame, exists := r.games[game.ID]
	if !exists {
		return fmt.Errorf("game with ID %s not found", game.ID)
	}

	// Preserve created timestamp, update modified timestamp
	game.CreatedAt = existingGame.CreatedAt
	game.UpdatedAt = time.Now().UTC()

	r.games[game.ID] = game
	return nil
}

// DeleteGame removes a game from the repository.
func (r *NBARepository) DeleteGame(gameID string) error {
	if gameID == "" {
		return fmt.Errorf("game ID is required")
	}

	if _, exists := r.games[gameID]; !exists {
		return fmt.Errorf("game with ID %s not found", gameID)
	}

	delete(r.games, gameID)
	return nil
}

// initializeSampleData creates some sample games for demonstration purposes.
func (r *NBARepository) initializeSampleData() {
	now := time.Now().UTC()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	// Sample teams
	lakers := models.Team{
		ID:           "lal",
		Name:         "Lakers",
		City:         "Los Angeles",
		Abbreviation: "LAL",
	}

	warriors := models.Team{
		ID:           "gsw",
		Name:         "Warriors",
		City:         "Golden State",
		Abbreviation: "GSW",
	}

	bulls := models.Team{
		ID:           "chi",
		Name:         "Bulls",
		City:         "Chicago",
		Abbreviation: "CHI",
	}

	// Sample arena
	staplesCenter := models.Arena{
		Name:     "Crypto.com Arena",
		City:     "Los Angeles",
		State:    "CA",
		Capacity: 20000,
	}

	// Sample games
	games := []*models.Game{
		{
			ID:        "game1",
			HomeTeam:  lakers,
			AwayTeam:  warriors,
			GameDate:  yesterday,
			Status:    string(models.GameStatusFinished),
			HomeScore: 108,
			AwayScore: 112,
			Period:    4,
			Arena:     staplesCenter,
			CreatedAt: yesterday,
			UpdatedAt: yesterday,
		},
		{
			ID:        "game2",
			HomeTeam:  bulls,
			AwayTeam:  lakers,
			GameDate:  now,
			Status:    string(models.GameStatusLive),
			HomeScore: 45,
			AwayScore: 52,
			Period:    2,
			GameClock: "8:23",
			Arena:     staplesCenter,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        "game3",
			HomeTeam:  warriors,
			AwayTeam:  bulls,
			GameDate:  tomorrow,
			Status:    string(models.GameStatusScheduled),
			HomeScore: 0,
			AwayScore: 0,
			Period:    0,
			Arena:     staplesCenter,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	// Add games to repository
	for _, game := range games {
		r.games[game.ID] = game
	}
}