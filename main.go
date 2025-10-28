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
		outputJSON = flag.String("json", "nba_results.json", "Output JSON file path")
		outputExcel = flag.String("excel", "nba_results.xlsx", "Output Excel file path")
		date = flag.String("date", time.Now().Format("2006-01-02"), "Date to fetch games for (YYYY-MM-DD)")
		help = flag.Bool("help", false, "Show help")
	)
	flag.Parse()

	if *help {
		fmt.Println("NBA Results Console App")
		fmt.Println("Usage:")
		fmt.Println("  -json string    Output JSON file path (default: nba_results.json)")
		fmt.Println("  -excel string   Output Excel file path (default: nba_results.xlsx)")
		fmt.Println("  -date string    Date to fetch games for YYYY-MM-DD (default: today)")
		fmt.Println("  -help           Show this help message")
		return
	}

	fmt.Printf("Fetching NBA game results for %s...\n", *date)

	// Create NBA client
	client := nba.NewClient()

	// Fetch games for the specified date
	games, err := client.GetGamesByDate(*date)
	if err != nil {
		log.Fatalf("Error fetching NBA games: %v", err)
	}

	if len(games) == 0 {
		fmt.Printf("No NBA games found for %s\n", *date)
		return
	}

	fmt.Printf("Found %d games\n", len(games))

	// Generate JSON output
	if err := generateJSONOutput(games, *outputJSON); err != nil {
		log.Fatalf("Error generating JSON output: %v", err)
	}

	// Generate Excel report
	if err := generateExcelReport(games, *outputExcel); err != nil {
		log.Fatalf("Error generating Excel report: %v", err)
	}

	fmt.Printf("Results saved to %s and %s\n", *outputJSON, *outputExcel)
}

func generateJSONOutput(games []nba.Game, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create JSON file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(games)
}

func generateExcelReport(games []nba.Game, filename string) error {
	reporter := report.NewExcelReporter()
	return reporter.GenerateReport(games, filename)
}