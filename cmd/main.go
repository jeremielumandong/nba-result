package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jeremielumandong/nba-result/internal/nba"
	"github.com/jeremielumandong/nba-result/internal/exporter"
)

func main() {
	// Initialize NBA client
	client := nba.NewClient()

	// Get today's games
	games, err := client.GetTodaysGames()
	if err != nil {
		log.Fatalf("Failed to fetch NBA games: %v", err)
	}

	// Generate JSON output
	jsonData, err := json.MarshalIndent(games, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Print JSON to console
	fmt.Println(string(jsonData))

	// Export to Excel
	excelExporter := exporter.NewExcelExporter()
	err = excelExporter.ExportGames(games, "nba_games_" + games[0].Date + ".xlsx")
	if err != nil {
		log.Printf("Warning: Failed to export to Excel: %v", err)
	} else {
		fmt.Printf("\nExcel report exported successfully to: nba_games_%s.xlsx\n", games[0].Date)
	}
}