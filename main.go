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
		startDate  = flag.String("start-date", "", "Start date for range query (YYYY-MM-DD)")
		endDate    = flag.String("end-date", "", "End date for range query (YYYY-MM-DD)")
		help       = flag.Bool("help", false, "Show help message")
	)
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	// Create NBA client and date service
	client := nba.NewClient()
	dateService := nba.NewDateService(client)

	// Handle date range query
	if *startDate != "" && *endDate != "" {
		handleDateRangeQuery(dateService, *startDate, *endDate, *outputFile, *excelFile)
		return
	}

	// Handle single date query
	var targetDateStr string
	if *date != "" {
		targetDateStr = *date
	} else {
		targetDateStr = time.Now().Format("2006-01-02")
	}

	handleSingleDateQuery(dateService, targetDateStr, *outputFile, *excelFile)
}

func handleSingleDateQuery(dateService *nba.DateService, dateStr, outputFile, excelFile string) {
	fmt.Printf("Fetching NBA games for %s...\n", dateStr)

	// Get games by date
	result, err := dateService.GetGamesByDate(dateStr)
	if err != nil {
		log.Fatalf("Error fetching NBA games: %v", err)
	}

	fmt.Printf("Found %d games\n", result.TotalGames)

	// Print summary
	printSummary(result.Summary)

	// Save JSON result
	if err := saveGameResultsJSON(result, outputFile); err != nil {
		log.Fatalf("Error saving JSON: %v", err)
	}
	fmt.Printf("JSON results saved to: %s\n", outputFile)

	// Generate Excel report
	reporter := report.NewExcelReporter()
	if err := reporter.GenerateReport(result.Games, excelFile); err != nil {
		log.Fatalf("Error generating Excel report: %v", err)
	}
	fmt.Printf("Excel report saved to: %s\n", excelFile)
}

func handleDateRangeQuery(dateService *nba.DateService, startDate, endDate, outputFile, excelFile string) {
	fmt.Printf("Fetching NBA games from %s to %s...\n", startDate, endDate)

	// Get games by date range
	results, err := dateService.GetGamesByDateRange(startDate, endDate)
	if err != nil {
		log.Fatalf("Error fetching NBA games for date range: %v", err)
	}

	// Aggregate all games and create summary
	allGames := []nba.Game{}
	totalGames := 0
	aggregatedSummary := nba.GameSummary{}

	for _, result := range results {
		allGames = append(allGames, result.Games...)
		totalGames += result.TotalGames
		aggregatedSummary.Final += result.Summary.Final
		aggregatedSummary.Live += result.Summary.Live
		aggregatedSummary.Scheduled += result.Summary.Scheduled
		aggregatedSummary.Other += result.Summary.Other
	}

	fmt.Printf("Found %d games across %d days\n", totalGames, len(results))
	printSummary(aggregatedSummary)

	// Create aggregated result for JSON export
	aggregatedResult := &nba.GameResults{
		Date:       fmt.Sprintf("%s to %s", startDate, endDate),
		Games:      allGames,
		TotalGames: totalGames,
		Summary:    aggregatedSummary,
		Metadata: nba.ResultMetadata{
			GeneratedAt: time.Now().Format(time.RFC3339),
			Source:      "NBA API",
			Version:     "1.0",
		},
	}

	// Save JSON result
	if err := saveGameResultsJSON(aggregatedResult, outputFile); err != nil {
		log.Fatalf("Error saving JSON: %v", err)
	}
	fmt.Printf("JSON results saved to: %s\n", outputFile)

	// Generate Excel report
	reporter := report.NewExcelReporter()
	if err := reporter.GenerateReport(allGames, excelFile); err != nil {
		log.Fatalf("Error generating Excel report: %v", err)
	}
	fmt.Printf("Excel report saved to: %s\n", excelFile)
}

func saveGameResultsJSON(result *nba.GameResults, filename string) error {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling JSON: %w", err)
	}

	return os.WriteFile(filename, data, 0644)
}

func printSummary(summary nba.GameSummary) {
	fmt.Println("\nGame Summary:")
	fmt.Printf("  Final: %d\n", summary.Final)
	fmt.Printf("  Live: %d\n", summary.Live)
	fmt.Printf("  Scheduled: %d\n", summary.Scheduled)
	if summary.Other > 0 {
		fmt.Printf("  Other: %d\n", summary.Other)
	}
	fmt.Println()
}

func printHelp() {
	fmt.Println("NBA Game Results Tracker")
	fmt.Println("========================")
	fmt.Println()
	fmt.Println("Usage: go run main.go [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -output string")
	fmt.Println("        Output JSON file path (default: nba_results.json)")
	fmt.Println("  -excel string")
	fmt.Println("        Output Excel file path (default: nba_results.xlsx)")
	fmt.Println("  -date string")
	fmt.Println("        Date in YYYY-MM-DD format (default: today)")
	fmt.Println("  -start-date string")
	fmt.Println("        Start date for range query (YYYY-MM-DD)")
	fmt.Println("  -end-date string")
	fmt.Println("        End date for range query (YYYY-MM-DD)")
	fmt.Println("  -help")
	fmt.Println("        Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run main.go                              # Get today's games")
	fmt.Println("  go run main.go -date 2024-01-15             # Get games for specific date")
	fmt.Println("  go run main.go -start-date 2024-01-15 -end-date 2024-01-17  # Get games for date range")
	fmt.Println("  go run main.go -output results.json         # Custom output file")
	fmt.Println("  go run main.go -excel report.xlsx           # Custom Excel file")
}