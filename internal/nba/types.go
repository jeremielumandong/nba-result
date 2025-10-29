package nba

// Game represents an NBA game
type Game struct {
	GameID   string `json:"game_id"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	HomeTeam Team   `json:"home_team"`
	AwayTeam Team   `json:"away_team"`
	Status   string `json:"status"` // "Scheduled", "Live", "Final"
	Quarter  int    `json:"quarter"`
	TimeLeft string `json:"time_left"`
}

// Team represents an NBA team
type Team struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Score int    `json:"score"`
}

// NBAAPIResponse represents the structure from NBA's API
// Note: This is a simplified structure. The actual NBA API has a more complex structure
type NBAAPIResponse struct {
	Scoreboard struct {
		Games []struct {
			GameID    string `json:"gameId"`
			GameCode  string `json:"gameCode"`
			GameStatus int    `json:"gameStatus"`
			HomeTeam  struct {
				TeamName string `json:"teamName"`
				TeamCode string `json:"teamTricode"`
				Score    int    `json:"score"`
			} `json:"homeTeam"`
			AwayTeam struct {
				TeamName string `json:"teamName"`
				TeamCode string `json:"teamTricode"`
				Score    int    `json:"score"`
			} `json:"awayTeam"`
		} `json:"games"`
	} `json:"scoreboard"`
}