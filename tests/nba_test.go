package tests

import (
	"testing"
	"time"

	"github.com/jeremielumandong/nba-result/internal/nba"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	client := nba.NewClient()
	assert.NotNil(t, client)
}

func TestGetGamesForDate(t *testing.T) {
	client := nba.NewClient()
	testDate := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)

	games, err := client.GetGamesForDate(testDate)
	
	require.NoError(t, err)
	assert.NotEmpty(t, games)
	
	// Verify game structure
	for _, game := range games {
		assert.NotEmpty(t, game.GameID)
		assert.NotEmpty(t, game.Date)
		assert.NotEmpty(t, game.HomeTeam.Name)
		assert.NotEmpty(t, game.AwayTeam.Name)
		assert.NotEmpty(t, game.HomeTeam.Code)
		assert.NotEmpty(t, game.AwayTeam.Code)
		assert.GreaterOrEqual(t, game.HomeTeam.Score, 0)
		assert.GreaterOrEqual(t, game.AwayTeam.Score, 0)
		assert.Contains(t, []string{"Scheduled", "Live", "Final"}, game.Status)
	}
}

func TestGetGamesForToday(t *testing.T) {
	client := nba.NewClient()
	today := time.Now()

	games, err := client.GetGamesForDate(today)
	
	require.NoError(t, err)
	// Should return at least mock data
	assert.NotEmpty(t, games)
}

func TestGameStructure(t *testing.T) {
	client := nba.NewClient()
	testDate := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)

	games, err := client.GetGamesForDate(testDate)
	require.NoError(t, err)
	require.NotEmpty(t, games)

	game := games[0]
	
	// Test required fields
	assert.NotEmpty(t, game.GameID, "GameID should not be empty")
	assert.NotEmpty(t, game.Date, "Date should not be empty")
	assert.NotEmpty(t, game.HomeTeam.Name, "Home team name should not be empty")
	assert.NotEmpty(t, game.AwayTeam.Name, "Away team name should not be empty")
	
	// Test team codes
	assert.Len(t, game.HomeTeam.Code, 3, "Team code should be 3 characters")
	assert.Len(t, game.AwayTeam.Code, 3, "Team code should be 3 characters")
	
	// Test scores are non-negative
	assert.GreaterOrEqual(t, game.HomeTeam.Score, 0, "Score should be non-negative")
	assert.GreaterOrEqual(t, game.AwayTeam.Score, 0, "Score should be non-negative")
	
	// Test status is valid
	validStatuses := []string{"Scheduled", "Live", "Final"}
	assert.Contains(t, validStatuses, game.Status, "Status should be valid")
	
	// Test quarter is valid (1-4, or 0 for not started)
	assert.GreaterOrEqual(t, game.Quarter, 0, "Quarter should be non-negative")
	assert.LessOrEqual(t, game.Quarter, 4, "Quarter should not exceed 4")
}

func TestDateFormatting(t *testing.T) {
	client := nba.NewClient()
	testDate := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)

	games, err := client.GetGamesForDate(testDate)
	require.NoError(t, err)
	require.NotEmpty(t, games)

	// Check that date is properly formatted
	for _, game := range games {
		_, err := time.Parse("2006-01-02", game.Date)
		assert.NoError(t, err, "Date should be in YYYY-MM-DD format")
	}
}