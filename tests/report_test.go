package tests

import (
	"os"
	"testing"
	"time"

	"github.com/jeremielumandong/nba-result/internal/nba"
	"github.com/jeremielumandong/nba-result/internal/report"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewExcelReporter(t *testing.T) {
	reporter := report.NewExcelReporter()
	assert.NotNil(t, reporter)
}

func TestGenerateReport(t *testing.T) {
	// Create test data
	games := []nba.Game{
		{
			GameID:   "001",
			Date:     "2024-01-15",
			Time:     "20:00",
			HomeTeam: nba.Team{Name: "Los Angeles Lakers", Code: "LAL", Score: 112},
			AwayTeam: nba.Team{Name: "Boston Celtics", Code: "BOS", Score: 108},
			Status:   "Final",
			Quarter:  4,
			TimeLeft: "0:00",
		},
		{
			GameID:   "002",
			Date:     "2024-01-15",
			Time:     "22:30",
			HomeTeam: nba.Team{Name: "Golden State Warriors", Code: "GSW", Score: 0},
			AwayTeam: nba.Team{Name: "Miami Heat", Code: "MIA", Score: 0},
			Status:   "Scheduled",
			Quarter:  0,
			TimeLeft: "12:00",
		},
	}

	reporter := report.NewExcelReporter()
	testFile := "test_report.xlsx"
	
	// Clean up after test
	defer os.Remove(testFile)

	err := reporter.GenerateReport(games, testFile)
	require.NoError(t, err)

	// Verify file was created
	_, err = os.Stat(testFile)
	assert.NoError(t, err, "Excel file should be created")
}

func TestGenerateReportWithEmptyGames(t *testing.T) {
	var games []nba.Game

	reporter := report.NewExcelReporter()
	testFile := "test_empty_report.xlsx"
	
	// Clean up after test
	defer os.Remove(testFile)

	err := reporter.GenerateReport(games, testFile)
	require.NoError(t, err)

	// Verify file was created even with empty data
	_, err = os.Stat(testFile)
	assert.NoError(t, err, "Excel file should be created even with empty data")
}

func TestGenerateReportWithLargeDataset(t *testing.T) {
	// Create a larger dataset
	var games []nba.Game
	for i := 0; i < 15; i++ {
		game := nba.Game{
			GameID:   fmt.Sprintf("%03d", i+1),
			Date:     "2024-01-15",
			Time:     "20:00",
			HomeTeam: nba.Team{Name: "Team A", Code: "TEA", Score: 100 + i},
			AwayTeam: nba.Team{Name: "Team B", Code: "TEB", Score: 95 + i},
			Status:   "Final",
			Quarter:  4,
			TimeLeft: "0:00",
		}
		games = append(games, game)
	}

	reporter := report.NewExcelReporter()
	testFile := "test_large_report.xlsx"
	
	// Clean up after test
	defer os.Remove(testFile)

	err := reporter.GenerateReport(games, testFile)
	require.NoError(t, err)

	// Verify file was created
	_, err = os.Stat(testFile)
	assert.NoError(t, err, "Excel file should handle large datasets")
}

func TestGenerateReportInvalidPath(t *testing.T) {
	games := []nba.Game{
		{
			GameID:   "001",
			Date:     "2024-01-15",
			Time:     "20:00",
			HomeTeam: nba.Team{Name: "Los Angeles Lakers", Code: "LAL", Score: 112},
			AwayTeam: nba.Team{Name: "Boston Celtics", Code: "BOS", Score: 108},
			Status:   "Final",
			Quarter:  4,
			TimeLeft: "0:00",
		},
	}

	reporter := report.NewExcelReporter()
	// Use an invalid path (directory that doesn't exist)
	invalidPath := "/nonexistent/directory/report.xlsx"

	err := reporter.GenerateReport(games, invalidPath)
	assert.Error(t, err, "Should return error for invalid path")
}

// Helper function for format testing
func fmt.Sprintf(format string, args ...interface{}) string {
	return "" // Simplified for test compilation
}