package nba

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// NBA API endpoint for today's games
	NBA_API_BASE_URL = "https://stats.nba.com/stats/scoreboardV2"
)

// Client handles NBA API requests
type Client struct {
	httpClient *http.Client
}

// Game represents an NBA game with results
type Game struct {
	GameID       string    `json:"game_id"`
	Date         string    `json:"date"`
	HomeTeam     Team      `json:"home_team"`
	AwayTeam     Team      `json:"away_team"`
	Status       string    `json:"status"`
	Period       int       `json:"period"`
	TimeRemaining string   `json:"time_remaining"`
	StartTime    time.Time `json:"start_time"`
}

// Team represents a basketball team
type Team struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Score        int    `json:"score"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
}

// APIResponse represents the NBA API response structure
type APIResponse struct {
	ResultSets []struct {
		Name    string        `json:"name"`
		Headers []string      `json:"headers"`
		RowSet  []interface{} `json:"rowSet"`
	} `json:"resultSets"`
}

// NewClient creates a new NBA API client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetTodaysGames fetches today's NBA games from the API
func (c *Client) GetTodaysGames() ([]Game, error) {
	today := time.Now().Format("20060102")
	url := fmt.Sprintf("%s?DayOffset=0&LeagueID=00&gameDate=%s", NBA_API_BASE_URL, today)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add required headers for NBA API
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://www.nba.com/")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResp APIResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return c.parseGames(apiResp, today)
}

// parseGames converts API response to Game structs
func (c *Client) parseGames(resp APIResponse, date string) ([]Game, error) {
	var games []Game

	// Find the GameHeader result set
	var gameHeaders []interface{}
	var lineScore []interface{}

	for _, resultSet := range resp.ResultSets {
		switch resultSet.Name {
		case "GameHeader":
			gameHeaders = resultSet.RowSet
		case "LineScore":
			lineScore = resultSet.RowSet
		}
	}

	if len(gameHeaders) == 0 {
		// No games today - return empty slice
		return games, nil
	}

	// Parse game data
	for _, gameData := range gameHeaders {
		gameRow, ok := gameData.([]interface{})
		if !ok || len(gameRow) < 11 {
			continue
		}

		game := Game{
			GameID: fmt.Sprintf("%v", gameRow[2]),
			Date:   date,
			Status: fmt.Sprintf("%v", gameRow[3]),
		}

		// Parse period and time remaining
		if gameRow[4] != nil {
			game.Period = int(gameRow[4].(float64))
		}
		if gameRow[5] != nil {
			game.TimeRemaining = fmt.Sprintf("%v", gameRow[5])
		}

		// Parse team data from LineScore
		c.parseTeamData(&game, lineScore)

		games = append(games, game)
	}

	return games, nil
}

// parseTeamData extracts team information from line score data
func (c *Client) parseTeamData(game *Game, lineScore []interface{}) {
	for _, scoreData := range lineScore {
		scoreRow, ok := scoreData.([]interface{})
		if !ok || len(scoreRow) < 8 {
			continue
		}

		gameID := fmt.Sprintf("%v", scoreRow[0])
		if gameID != game.GameID {
			continue
		}

		team := Team{
			ID:           int(scoreRow[1].(float64)),
			Abbreviation: fmt.Sprintf("%v", scoreRow[2]),
			Name:         fmt.Sprintf("%v", scoreRow[3]),
		}

		// Parse score
		if scoreRow[7] != nil {
			team.Score = int(scoreRow[7].(float64))
		}

		// Determine if home or away team
		if scoreRow[4] != nil && scoreRow[4].(string) == "1" {
			game.HomeTeam = team
		} else {
			game.AwayTeam = team
		}
	}
}