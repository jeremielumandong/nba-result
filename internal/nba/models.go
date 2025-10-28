package nba

import "time"

// Game represents an NBA game with all relevant information
type Game struct {
	GameID       string    `json:"game_id"`
	GameDate     string    `json:"game_date"`
	GameTime     string    `json:"game_time"`
	HomeTeam     Team      `json:"home_team"`
	AwayTeam     Team      `json:"away_team"`
	Status       string    `json:"status"`
	SeasonType   string    `json:"season_type"`
	CreatedAt    time.Time `json:"created_at"`
}

// Team represents a basketball team
type Team struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// Winner returns the winning team name, or "TIE" if tied
func (g *Game) Winner() string {
	if g.HomeTeam.Score > g.AwayTeam.Score {
		return g.HomeTeam.Name
	} else if g.AwayTeam.Score > g.HomeTeam.Score {
		return g.AwayTeam.Name
	}
	return "TIE"
}

// IsFinished returns true if the game is completed
func (g *Game) IsFinished() bool {
	return g.Status == "Final" || g.Status == "3 - Final"
}

// APIResponse represents the structure of the NBA API response
type APIResponse struct {
	Resource   string      `json:"resource"`
	Parameters interface{} `json:"parameters"`
	ResultSets []ResultSet `json:"resultSets"`
}

// ResultSet represents a result set from the NBA API
type ResultSet struct {
	Name    string          `json:"name"`
	Headers []string        `json:"headers"`
	RowSet  [][]interface{} `json:"rowSet"`
}

// GameSummary provides a brief summary of the game
type GameSummary struct {
	Matchup    string `json:"matchup"`
	FinalScore string `json:"final_score"`
	Winner     string `json:"winner"`
	Status     string `json:"status"`
	Date       string `json:"date"`
}

// ToSummary converts a Game to a GameSummary
func (g *Game) ToSummary() GameSummary {
	return GameSummary{
		Matchup:    fmt.Sprintf("%s vs %s", g.AwayTeam.Name, g.HomeTeam.Name),
		FinalScore: fmt.Sprintf("%d - %d", g.AwayTeam.Score, g.HomeTeam.Score),
		Winner:     g.Winner(),
		Status:     g.Status,
		Date:       g.GameDate,
	}
}