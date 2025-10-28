package nba

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// Using NBA API endpoint (unofficial but commonly used)
	baseURL = "https://data.nba.net/10s/prod/v1"
	userAgent = "NBA-Result-Tracker/1.0"
)

// Client represents the NBA API client
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

// GetGamesByDate fetches games for a specific date
func (c *Client) GetGamesByDate(date time.Time) ([]Game, error) {
	dateStr := date.Format("20060102")
	url := fmt.Sprintf("%s/%s/scoreboard.json", baseURL, dateStr)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response ScoreboardResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Convert API response to our Game structure
	games := make([]Game, len(response.Games))
	for i, apiGame := range response.Games {
		games[i] = convertAPIGameToGame(apiGame, date)
	}

	return games, nil
}

// convertAPIGameToGame converts API response game to our Game struct
func convertAPIGameToGame(apiGame APIGame, date time.Time) Game {
	game := Game{
		GameID: apiGame.GameID,
		Date: date.Format("2006-01-02"),
		HomeTeam: Team{
			TeamID: apiGame.HTeam.TeamID,
			Tricode: apiGame.HTeam.Tricode,
			Score: apiGame.HTeam.Score,
		},
		VisitorTeam: Team{
			TeamID: apiGame.VTeam.TeamID,
			Tricode: apiGame.VTeam.Tricode,
			Score: apiGame.VTeam.Score,
		},
		Period: apiGame.Period.Current,
		Clock: apiGame.Clock,
		StatusText: apiGame.StatusText,
	}

	// Determine game status
	if apiGame.IsGameActivated {
		if apiGame.Period.Current == 0 {
			game.Status = "Scheduled"
		} else if apiGame.Period.Current > 0 && apiGame.Period.Current <= 4 {
			game.Status = "Live"
		} else {
			game.Status = "Final"
		}
	} else {
		game.Status = "Final"
	}

	// Determine winner
	if game.Status == "Final" {
		homeScore := parseScore(apiGame.HTeam.Score)
		visitorScore := parseScore(apiGame.VTeam.Score)
		if homeScore > visitorScore {
			game.Winner = "Home"
		} else if visitorScore > homeScore {
			game.Winner = "Visitor"
		} else {
			game.Winner = "Tie"
		}
	}

	return game
}

// parseScore safely parses score string to int
func parseScore(scoreStr string) int {
	if scoreStr == "" {
		return 0
	}
	var score int
	fmt.Sscanf(scoreStr, "%d", &score)
	return score
}