package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jeremielumandong/nba-result/internal/nba"
	"github.com/jeremielumandong/nba-result/internal/exporter"
)

func main() {
	// Command line flags
	var (
		jsonFile = flag.String("json", "nba_games.json", "Output JSON file path")
		excelFile = flag.String("excel", "nba_games.xlsx", "Output Excel file path")
		date = flag.String("date", "", "Date to fetch games for (YYYY-MM-DD format, defaults to today)")
		help = flag.Bool("help", false, "Show help message")
	)
	flag.Parse()

	if *help {
		fmt.Printf("NBA Game Results Tracker\n\n")
		fmt.Printf("Usage: %s [options]\n\n", os.Args[0])
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
		return
	}

	// Parse date or use today
	var targetDate time.Time
	var err error
	if *date == "" {
		targetDate = time.Now()
	} else {
		targetDate, err = time.Parse("2006-01-02", *date)
		if err != nil {
			log.Fatalf("Invalid date format. Use YYYY-MM-DD: %v", err)
		}
	}

	fmt.Printf("Fetching NBA games for %s...\n", targetDate.Format("2006-01-02"))

	// Create NBA client
	client := nba.NewClient()

	// Fetch games
	games, err := client.GetGamesByDate(targetDate)
	if err != nil {
		log.Fatalf("Failed to fetch games: %v", err)
	}

	fmt.Printf("Found %d games\n", len(games))

	// Export to JSON
	if err := exporter.ExportToJSON(games, *jsonFile); err != nil {
		log.Fatalf("Failed to export JSON: %v", err)
	}
	fmt.Printf("JSON exported to: %s\n", *jsonFile)

	// Export to Excel
	if err := exporter.ExportToExcel(games, *excelFile); err != nil {
		log.Fatalf("Failed to export Excel: %v", err)
	}
	fmt.Printf("Excel exported to: %s\n", *excelFile)

	fmt.Println("Export completed successfully!")
}