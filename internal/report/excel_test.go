package report

import (
	"os"
	"testing"
	"time"

	"github.com/jeremielumandong/nba-result/internal/nba"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewExcelReporter(t *testing.T) {
	reporter := NewExcelReporter()
	assert.NotNil(t, reporter)
	assert.NotNil(t, reporter.file)
}

func TestExcelReporter_GenerateReport(t *testing.T) {
	// Create test data
	games := []nba.Game{
		{
			GameID:     "0022300567",
			GameDate:   "2024-01-15",
			GameTime:   "19:00",
			HomeTeam:   nba.Team{Name: "Lakers", Score: 110},
			AwayTeam:   nba.Team{Name: "Warriors", Score: 105},
			Status:     "Final",
			SeasonType: "Regular Season",
			CreatedAt:  time.Now(),
		},
		{
			GameID:     "0022300568",
			GameDate:   "2024-01-15",
			GameTime:   "21:30",
			HomeTeam:   nba.Team{Name: "Celtics", Score: 95},
			AwayTeam:   nba.Team{Name: "Heat", Score: 88},
			Status:     "2nd Qtr",
			SeasonType: "Regular Season",
			CreatedAt:  time.Now(),
		},
	}

	reporter := NewExcelReporter()
	filename := "test_nba_results.xlsx"

	// Clean up after test
	defer func() {
		os.Remove(filename)
	}()

	// Generate report
	err := reporter.GenerateReport(games, filename)
	require.NoError(t, err)

	// Verify file was created
	_, err = os.Stat(filename)
	assert.NoError(t, err)
}

func TestExcelReporter_GenerateReport_EmptyGames(t *testing.T) {
	reporter := NewExcelReporter()
	filename := "test_empty_results.xlsx"

	// Clean up after test
	defer func() {
		os.Remove(filename)
	}()

	// Generate report with empty games
	err := reporter.GenerateReport([]nba.Game{}, filename)
	require.NoError(t, err)

	// Verify file was created
	_, err = os.Stat(filename)
	assert.NoError(t, err)
}

func TestExcelReporter_GenerateReport_InvalidPath(t *testing.T) {
	reporter := NewExcelReporter()
	games := []nba.Game{
		{
			GameID:   "test",
			HomeTeam: nba.Team{Name: "Lakers", Score: 110},
			AwayTeam: nba.Team{Name: "Warriors", Score: 105},
		},
	}

	// Try to save to invalid path
	err := reporter.GenerateReport(games, "/invalid/path/file.xlsx")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to save Excel file")
}