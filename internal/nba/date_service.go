package nba

import (
	"fmt"
	"time"
)

// DateService handles date-related operations for NBA games
type DateService struct {
	client *Client
}

// NewDateService creates a new DateService
func NewDateService(client *Client) *DateService {
	return &DateService{
		client: client,
	}
}

// GetGamesByDate fetches NBA games for a specific date and returns structured results
func (ds *DateService) GetGamesByDate(dateStr string) (*GameResults, error) {
	// Parse the date string
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format '%s': use YYYY-MM-DD format: %w", dateStr, err)
	}

	// Validate date is not in the future
	if date.After(time.Now()) {
		return nil, fmt.Errorf("date cannot be in the future")
	}

	// Validate date is not too far in the past (NBA founded in 1946)
	nbaFoundedYear := 1946
	if date.Year() < nbaFoundedYear {
		return nil, fmt.Errorf("date cannot be before NBA was founded (%d)", nbaFoundedYear)
	}

	// Fetch games for the date
	games, err := ds.client.GetGamesForDate(date)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch games for date %s: %w", dateStr, err)
	}

	// Create structured result
	result := &GameResults{
		Date:       dateStr,
		Games:      games,
		TotalGames: len(games),
		Summary:    ds.generateSummary(games),
		Metadata: ResultMetadata{
			GeneratedAt: time.Now().Format(time.RFC3339),
			Source:      "NBA API",
			Version:     "1.0",
		},
	}

	return result, nil
}

// generateSummary creates a summary of game statuses
func (ds *DateService) generateSummary(games []Game) GameSummary {
	summary := GameSummary{}
	
	for _, game := range games {
		switch game.Status {
		case "Scheduled":
			summary.Scheduled++
		case "Live":
			summary.Live++
		case "Final":
			summary.Final++
		default:
			summary.Other++
		}
	}
	
	return summary
}

// GetGamesByDateRange fetches NBA games for a date range
func (ds *DateService) GetGamesByDateRange(startDateStr, endDateStr string) ([]*GameResults, error) {
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format '%s': use YYYY-MM-DD format: %w", startDateStr, err)
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format '%s': use YYYY-MM-DD format: %w", endDateStr, err)
	}

	if endDate.Before(startDate) {
		return nil, fmt.Errorf("end date cannot be before start date")
	}

	// Limit range to prevent excessive API calls (max 30 days)
	maxDays := 30
	daysDiff := int(endDate.Sub(startDate).Hours() / 24)
	if daysDiff > maxDays {
		return nil, fmt.Errorf("date range too large: maximum %d days allowed", maxDays)
	}

	var results []*GameResults
	currentDate := startDate

	for !currentDate.After(endDate) {
		dateStr := currentDate.Format("2006-01-02")
		result, err := ds.GetGamesByDate(dateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to get games for %s: %w", dateStr, err)
		}
		results = append(results, result)
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return results, nil
}