package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jeremielumandong/nba-result/internal/nba"
	"github.com/jeremielumandong/nba-result/internal/report"
)

func main() {
	// Command line flags
	var (
		outputFile = flag.String("output", "nba_results.json", "Output JSON file path")
		excelFile  = flag.String("excel", "nba_results.xlsx", "Output Excel file path")
		date       = flag.String("date", "", "Date in YYYY-MM-DD format (default: today)")
		help       = flag.Bool("help", false, "Show help message")
	)
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	// Parse date or use today
	targetDate := time.Now()
	if *date != "" {
		parsedDate, err := time.Parse("2006-01-02", *date)
		if err != nil {
			log.Fatalf("Invalid date format. Use YYYY-MM-DD: %v", err)
		}
		targetDate = parsedDate
	}

	fmt.Printf("Fetching NBA games for %s...\n", targetDate.Format("2006-01-02"))

	// Create NBA client
	client := nba.NewClient()

	// Fetch games
	games, err := client.GetGamesForDate(targetDate)
	if err != nil {
		log.Fatalf("Error fetching NBA games: %v", err)
	}

	fmt.Printf("Found %d games\n", len(games))

	// Save JSON result
	if err := saveJSON(games, *outputFile); err != nil {
		log.Fatalf("Error saving JSON: %v", err)
	}
	fmt.Printf("JSON results saved to: %s\n", *outputFile)

	// Generate Excel report
	reporter := report.NewExcelReporter()
	if err := reporter.GenerateReport(games, *excelFile); err != nil {
		log.Fatalf("Error generating Excel report: %v", err)
	}
	fmt.Printf("Excel report saved to: %s\n", *excelFile)
}

func saveJSON(games []nba.Game, filename string) error {
	data, err := json.MarshalIndent(games, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling JSON: %w", err)
	}

	return os.WriteFile(filename, data, 0644)
}

func printHelp() {
	fmt.Println("NBA Game Results Tracker")
	fmt.Println("========================")
	fmt.Println()
	fmt.Println("Usage: go run main.go [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -output string    Output JSON file path (default: nba_results.json)")
	fmt.Println("  -excel string     Output Excel file path (default: nba_results.xlsx)")
	fmt.Println("  -date string      Date in YYYY-MM-DD format (default: today)")
	fmt.Println("  -help             Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run main.go")
	fmt.Println("  go run main.go -date 2024-01-15")
	fmt.Println("  go run main.go -output results.json -excel report.xlsx")
}