package exporter

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jeremielumandong/nba-result/internal/nba"
)

// GameResult represents the formatted game result for export
type GameResult struct {
	ExportInfo ExportInfo  `json:"export_info"`
	Games      []nba.Game  `json:"games"`
	Summary    GameSummary `json:"summary"`
}

// ExportInfo contains metadata about the export
type ExportInfo struct {
	GeneratedAt string `json:"generated_at"`
	TotalGames  int    `json:"total_games"`
	Date        string `json:"date"`
}

// GameSummary provides summary statistics
type GameSummary struct {
	Scheduled int `json:"scheduled"`
	Live      int `json:"live"`
	Final     int `json:"final"`
}

// ExportToJSON exports games to a JSON file
func ExportToJSON(games []nba.Game, filename string) error {
	if len(games) == 0 {
		return fmt.Errorf("no games to export")
	}

	// Create summary
	summary := GameSummary{}
	for _, game := range games {
		switch game.Status {
		case "Scheduled":
			summary.Scheduled++
		case "Live":
			summary.Live++
		case "Final":
			summary.Final++
		}
	}

	// Create result structure
	result := GameResult{
		ExportInfo: ExportInfo{
			GeneratedAt: time.Now().Format(time.RFC3339),
			TotalGames:  len(games),
			Date:        games[0].Date,
		},
		Games:   games,
		Summary: summary,
	}

	// Marshal to JSON with proper formatting
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Write to file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write JSON data: %w", err)
	}

	return nil
}