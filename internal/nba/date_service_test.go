package nba

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDateService(t *testing.T) {
	client := NewClient()
	dateService := NewDateService(client)
	assert.NotNil(t, dateService)
	assert.NotNil(t, dateService.client)
}

func TestGetGamesByDate_ValidDate(t *testing.T) {
	client := NewClient()
	dateService := NewDateService(client)

	// Test with a valid past date
	dateStr := "2024-01-15"
	result, err := dateService.GetGamesByDate(dateStr)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, dateStr, result.Date)
	assert.GreaterOrEqual(t, result.TotalGames, 0)
	assert.Equal(t, len(result.Games), result.TotalGames)
	assert.NotEmpty(t, result.Metadata.GeneratedAt)
	assert.Equal(t, "NBA API", result.Metadata.Source)
}

func TestGetGamesByDate_InvalidDateFormat(t *testing.T) {
	client := NewClient()
	dateService := NewDateService(client)

	testCases := []string{
		"2024/01/15", // Wrong format
		"01-15-2024", // Wrong format
		"2024-1-15",  // Missing leading zero
		"invalid",    // Not a date
		"",           // Empty string
	}

	for _, dateStr := range testCases {
		t.Run(dateStr, func(t *testing.T) {
			result, err := dateService.GetGamesByDate(dateStr)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "invalid date format")
		})
	}
}

func TestGetGamesByDate_FutureDate(t *testing.T) {
	client := NewClient()
	dateService := NewDateService(client)

	// Test with future date
	futureDate := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	result, err := dateService.GetGamesByDate(futureDate)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "date cannot be in the future")
}

func TestGetGamesByDate_PreNBADate(t *testing.T) {
	client := NewClient()
	dateService := NewDateService(client)

	// Test with date before NBA was founded
	result, err := dateService.GetGamesByDate("1945-01-01")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "date cannot be before NBA was founded")
}

func TestGenerateSummary(t *testing.T) {
	client := NewClient()
	dateService := NewDateService(client)

	// Create test games with different statuses
	games := []Game{
		{Status: "Final"},
		{Status: "Final"},
		{Status: "Live"},
		{Status: "Scheduled"},
		{Status: "Scheduled"},
		{Status: "Unknown"},
	}

	summary := dateService.generateSummary(games)

	assert.Equal(t, 2, summary.Final)
	assert.Equal(t, 1, summary.Live)
	assert.Equal(t, 2, summary.Scheduled)
	assert.Equal(t, 1, summary.Other)
}

func TestGetGamesByDateRange_ValidRange(t *testing.T) {
	client := NewClient()
	dateService := NewDateService(client)

	// Test with a small valid range
	startDate := "2024-01-15"
	endDate := "2024-01-17"

	results, err := dateService.GetGamesByDateRange(startDate, endDate)

	require.NoError(t, err)
	assert.NotNil(t, results)
	assert.Len(t, results, 3) // 3 days inclusive

	// Verify dates are correct
	expectedDates := []string{"2024-01-15", "2024-01-16", "2024-01-17"}
	for i, result := range results {
		assert.Equal(t, expectedDates[i], result.Date)
	}
}

func TestGetGamesByDateRange_InvalidDates(t *testing.T) {
	client := NewClient()
	dateService := NewDateService(client)

	// Test with invalid start date
	results, err := dateService.GetGamesByDateRange("invalid", "2024-01-17")
	assert.Error(t, err)
	assert.Nil(t, results)

	// Test with invalid end date
	results, err = dateService.GetGamesByDateRange("2024-01-15", "invalid")
	assert.Error(t, err)
	assert.Nil(t, results)
}

func TestGetGamesByDateRange_EndBeforeStart(t *testing.T) {
	client := NewClient()
	dateService := NewDateService(client)

	// Test with end date before start date
	results, err := dateService.GetGamesByDateRange("2024-01-17", "2024-01-15")
	assert.Error(t, err)
	assert.Nil(t, results)
	assert.Contains(t, err.Error(), "end date cannot be before start date")
}

func TestGetGamesByDateRange_TooLargeRange(t *testing.T) {
	client := NewClient()
	dateService := NewDateService(client)

	// Test with range larger than 30 days
	startDate := "2024-01-01"
	endDate := "2024-02-15" // More than 30 days

	results, err := dateService.GetGamesByDateRange(startDate, endDate)
	assert.Error(t, err)
	assert.Nil(t, results)
	assert.Contains(t, err.Error(), "date range too large")
}