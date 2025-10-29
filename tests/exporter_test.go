package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jeremielumandong/nba-result/internal/exporter"
	"github.com/jeremielumandong/nba-result/internal/nba"
	"github.com/xuri/excelize/v2"
)

func TestExportToJSON(t *testing.T) {
	// Create test data
	testGames := []nba.Game{
		{
			GameID: "0022301234",
			Date:   "2024-01-15",
			HomeTeam: nba.Team{
				TeamID:  "1610612747",
				Tricode: "LAL",
				Score:   "110",
			},
			VisitorTeam: nba.Team{
				TeamID:  "1610612744",
				Tricode: "GSW",
				Score:   "105",
			},
			Period:     4,
			Clock:      "",
			Status:     "Final",
			StatusText: "Final",
			Winner:     "Home",
		},
		{
			GameID: "0022301235",
			Date:   "2024-01-15",
			HomeTeam: nba.Team{
				TeamID:  "1610612738",
				Tricode: "BOS",
				Score:   "95",
			},
			VisitorTeam: nba.Team{
				TeamID:  "1610612752",
				Tricode: "NYK",
				Score:   "100",
			},
			Period:     4,
			Clock:      "",
			Status:     "Final",
			StatusText: "Final",
			Winner:     "Visitor",
		},
	}

	// Create temporary file
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "test_games.json")

	// Test export
	err := exporter.ExportToJSON(testGames, filename)
	if err != nil {
		t.Fatalf("ExportToJSON failed: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatal("JSON file was not created")
	}

	// Read and validate JSON content
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}

	var result exporter.GameResult
	if err := json.Unmarshal(content, &result); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Validate structure
	if result.ExportInfo.TotalGames != 2 {
		t.Errorf("Expected 2 games, got %d", result.ExportInfo.TotalGames)
	}
	if len(result.Games) != 2 {
		t.Errorf("Expected 2 games in array, got %d", len(result.Games))
	}
	if result.Summary.Final != 2 {
		t.Errorf("Expected 2 final games, got %d", result.Summary.Final)
	}
	if result.ExportInfo.Date != "2024-01-15" {
		t.Errorf("Expected date 2024-01-15, got %s", result.ExportInfo.Date)
	}
}

func TestExportToExcel(t *testing.T) {
	// Create test data
	testGames := []nba.Game{
		{
			GameID: "0022301234",
			Date:   "2024-01-15",
			HomeTeam: nba.Team{
				TeamID:  "1610612747",
				Tricode: "LAL",
				Score:   "110",
			},
			VisitorTeam: nba.Team{
				TeamID:  "1610612744",
				Tricode: "GSW",
				Score:   "105",
			},
			Period:     4,
			Clock:      "",
			Status:     "Final",
			StatusText: "Final",
			Winner:     "Home",
		},
		{
			GameID: "0022301235",
			Date:   "2024-01-15",
			HomeTeam: nba.Team{
				TeamID:  "1610612738",
				Tricode: "BOS",
				Score:   "0",
			},
			VisitorTeam: nba.Team{
				TeamID:  "1610612752",
				Tricode: "NYK",
				Score:   "0",
			},
			Period:     0,
			Clock:      "12:00",
			Status:     "Scheduled",
			StatusText: "7:30 pm ET",
			Winner:     "",
		},
	}

	// Create temporary file
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "test_games.xlsx")

	// Test export
	err := exporter.ExportToExcel(testGames, filename)
	if err != nil {
		t.Fatalf("ExportToExcel failed: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatal("Excel file was not created")
	}

	// Open and validate Excel file
	f, err := excelize.OpenFile(filename)
	if err != nil {
		t.Fatalf("Failed to open Excel file: %v", err)
	}
	defer f.Close()

	// Check sheets exist
	sheets := f.GetSheetList()
	expectedSheets := []string{"NBA Games", "Summary"}
	for _, expected := range expectedSheets {
		found := false
		for _, sheet := range sheets {
			if sheet == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected sheet '%s' not found", expected)
		}
	}

	// Check main sheet content
	mainSheet := "NBA Games"
	rows, err := f.GetRows(mainSheet)
	if err != nil {
		t.Fatalf("Failed to get rows from main sheet: %v", err)
	}

	// Should have header + 2 data rows
	if len(rows) != 3 {
		t.Errorf("Expected 3 rows (1 header + 2 data), got %d", len(rows))
	}

	// Check header row
	if len(rows) > 0 {
		headerRow := rows[0]
		expectedHeaders := []string{"Game ID", "Date", "Home Team", "Home Score", "Visitor Team", "Visitor Score", "Period", "Clock", "Status", "Winner"}
		for i, expected := range expectedHeaders {
			if i < len(headerRow) && headerRow[i] != expected {
				t.Errorf("Expected header '%s', got '%s'", expected, headerRow[i])
			}
		}
	}

	// Check summary sheet
	summaryRows, err := f.GetRows("Summary")
	if err != nil {
		t.Fatalf("Failed to get rows from summary sheet: %v", err)
	}

	if len(summaryRows) == 0 {
		t.Error("Summary sheet should not be empty")
	}
}

func TestExportEmptyGames(t *testing.T) {
	// Test with empty games slice
	emptyGames := []nba.Game{}
	tempDir := t.TempDir()

	// JSON export
	jsonFilename := filepath.Join(tempDir, "empty.json")
	err := exporter.ExportToJSON(emptyGames, jsonFilename)
	if err == nil {
		t.Error("Expected error when exporting empty games to JSON")
	}

	// Excel export
	excelFilename := filepath.Join(tempDir, "empty.xlsx")
	err = exporter.ExportToExcel(emptyGames, excelFilename)
	if err == nil {
		t.Error("Expected error when exporting empty games to Excel")
	}
}

func TestGameSummaryCalculation(t *testing.T) {
	// Test games with different statuses
	testGames := []nba.Game{
		{Status: "Final"},
		{Status: "Final"},
		{Status: "Live"},
		{Status: "Scheduled"},
		{Status: "Scheduled"},
		{Status: "Scheduled"},
	}

	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "summary_test.json")

	// Set a consistent date for all games
	for i := range testGames {
		testGames[i].Date = "2024-01-15"
		testGames[i].GameID = "test" + string(rune(i))
	}

	err := exporter.ExportToJSON(testGames, filename)
	if err != nil {
		t.Fatalf("Failed to export JSON: %v", err)
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	var result exporter.GameResult
	if err := json.Unmarshal(content, &result); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Check summary counts
	if result.Summary.Final != 2 {
		t.Errorf("Expected 2 final games, got %d", result.Summary.Final)
	}
	if result.Summary.Live != 1 {
		t.Errorf("Expected 1 live game, got %d", result.Summary.Live)
	}
	if result.Summary.Scheduled != 3 {
		t.Errorf("Expected 3 scheduled games, got %d", result.Summary.Scheduled)
	}
}