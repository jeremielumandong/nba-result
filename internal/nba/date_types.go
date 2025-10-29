package nba

// GameResults represents the structured result for games on a specific date
type GameResults struct {
	Date       string         `json:"date"`
	Games      []Game         `json:"games"`
	TotalGames int            `json:"total_games"`
	Summary    GameSummary    `json:"summary"`
	Metadata   ResultMetadata `json:"metadata"`
}

// GameSummary provides statistics about games
type GameSummary struct {
	Scheduled int `json:"scheduled"`
	Live      int `json:"live"`
	Final     int `json:"final"`
	Other     int `json:"other"`
}

// ResultMetadata contains metadata about the query result
type ResultMetadata struct {
	GeneratedAt string `json:"generated_at"`
	Source      string `json:"source"`
	Version     string `json:"version"`
}

// DateQueryRequest represents a request for games by date
type DateQueryRequest struct {
	Date       string `json:"date"`
	OutputJSON string `json:"output_json,omitempty"`
	OutputExcel string `json:"output_excel,omitempty"`
}

// DateRangeQueryRequest represents a request for games by date range
type DateRangeQueryRequest struct {
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	OutputJSON  string `json:"output_json,omitempty"`
	OutputExcel string `json:"output_excel,omitempty"`
}