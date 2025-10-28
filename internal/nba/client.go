// Package nba provides functionality to fetch NBA game data from external APIs
package nba

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	// NBA API base URL - using NBA official stats API
	baseURL = "https://stats.nba.com/stats/scoreboardV2"
	userAgent = "NBA-Results-App/1.0"
)

// Client represents an NBA API client
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new NBA API client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetGamesByDate fetches NBA games for a specific date
func (c *Client) GetGamesByDate(date string) ([]Game, error) {
	// Parse and validate date
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}

	// Format date for API
	gameDate := parsedDate.Format("01/02/2006")

	// Build request URL
	url := fmt.Sprintf("%s?GameDate=%s&LeagueID=00&DayOffset=0", baseURL, gameDate)

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers (required by NBA API)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://www.nba.com/")
	req.Header.Set("Origin", "https://www.nba.com")

	// Make request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Parse response
	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert API response to our Game struct
	games := c.parseGames(&apiResponse)
	return games, nil
}

// parseGames converts API response to Game structs
func (c *Client) parseGames(apiResp *APIResponse) []Game {
	var games []Game

	if len(apiResp.ResultSets) == 0 {
		return games
	}

	// Find the GameHeader result set
	var gameHeaders [][]interface{}
	var lineScores [][]interface{}

	for _, rs := range apiResp.ResultSets {
		if rs.Name == "GameHeader" {
			gameHeaders = rs.RowSet
		} else if rs.Name == "LineScore" {
			lineScores = rs.RowSet
		}
	}

	// Parse game headers
	for _, row := range gameHeaders {
		if len(row) < 12 {
			continue
		}

		game := Game{
			GameID:       toString(row[2]),
			GameDate:     toString(row[0]),
			HomeTeam:     Team{Name: toString(row[6])},
			AwayTeam:     Team{Name: toString(row[7])},
			Status:       toString(row[3]),
			SeasonType:   toString(row[4]),
			GameTime:     toString(row[9]),
		}

		// Find corresponding line scores for this game
		for _, lineRow := range lineScores {
			if len(lineRow) < 22 {
				continue
			}

			gameID := toString(lineRow[2])
			if gameID != game.GameID {
				continue
			}

			teamName := toString(lineRow[4])
			pts := toInt(lineRow[21])

			if teamName == game.HomeTeam.Name {
				game.HomeTeam.Score = pts
			} else if teamName == game.AwayTeam.Name {
				game.AwayTeam.Score = pts
			}
		}

		games = append(games, game)
	}

	return games
}

// Helper functions for type conversion
func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", v)
}

func toInt(v interface{}) int {
	if v == nil {
		return 0
	}
	switch val := v.(type) {
	case int:
		return val
	case float64:
		return int(val)
	case string:
		if val == "" {
			return 0
		}
		// Try to parse as number
		return 0
	default:
		return 0
	}
}