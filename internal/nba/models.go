package nba

// Game represents an NBA game with results
type Game struct {
	GameID      string `json:"game_id"`
	Date        string `json:"date"`
	HomeTeam    Team   `json:"home_team"`
	VisitorTeam Team   `json:"visitor_team"`
	Period      int    `json:"period"`
	Clock       string `json:"clock"`
	Status      string `json:"status"` // Scheduled, Live, Final
	StatusText  string `json:"status_text"`
	Winner      string `json:"winner,omitempty"` // Home, Visitor, or empty if not final
}

// Team represents a basketball team
type Team struct {
	TeamID  string `json:"team_id"`
	Tricode string `json:"tricode"` // Team abbreviation (e.g., LAL, GSW)
	Score   string `json:"score"`
}

// ScoreboardResponse represents the API response from NBA scoreboard endpoint
type ScoreboardResponse struct {
	Internal struct {
		PubDateTime string `json:"pubDateTime"`
	} `json:"internal"`
	Games []APIGame `json:"games"`
}

// APIGame represents a game from the NBA API
type APIGame struct {
	GameID           string `json:"gameId"`
	IsGameActivated  bool   `json:"isGameActivated"`
	StatusText       string `json:"statusText"`
	Clock            string `json:"clock"`
	IsHalftime       bool   `json:"isHalftime"`
	IsEndOfPeriod    bool   `json:"isEndOfPeriod"`
	HTeam           APITeam `json:"hTeam"`
	VTeam           APITeam `json:"vTeam"`
	Period          Period  `json:"period"`
	Nugget          struct {
		Text string `json:"text"`
	} `json:"nugget"`
}

// APITeam represents a team from the NBA API
type APITeam struct {
	TeamID   string `json:"teamId"`
	Tricode  string `json:"tricode"`
	Score    string `json:"score"`
	Win      string `json:"win"`
	Loss     string `json:"loss"`
	SeriesWin string `json:"seriesWin"`
	SeriesLoss string `json:"seriesLoss"`
}

// Period represents the game period information
type Period struct {
	Current    int  `json:"current"`
	Type       int  `json:"type"`
	MaxRegular int  `json:"maxRegular"`
	IsHalftime bool `json:"isHalftime"`
	IsEndOfPeriod bool `json:"isEndOfPeriod"`
}