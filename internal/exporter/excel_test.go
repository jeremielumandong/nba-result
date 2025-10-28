package exporter

import (
	"os"
	"testing"
	"time"

	"github.com/jeremielumandong/nba-result/internal/nba"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuri/excelize/v2"
)

func TestNewExcelExporter(t *testing.T) {
	exporter := NewExcelExporter()
	assert.NotNil(t, exporter)
}

func TestExportGames_Success(t *testing.T) {
	exporter := NewExcelExporter()

	// Create test data
	games := []nba.Game{
		{
			GameID:        "0022300123",
			Date:          "20240115",
			Status:        "Final",
			Period:        4,
			TimeRemaining: "",
			HomeTeam: nba.Team{
				ID:           1,
				Name:         "Los Angeles Lakers",
				Abbreviation: "LAL",
				Score:        110,
			},
			AwayTeam: nba.Team{
				ID:           2,
				Name:         "Golden State Warriors",
				Abbreviation: "GSW",
				Score:        105,
			},
			StartTime: time.Now(),
		},
		{
			GameID:        "0022300124",
			Date:          "20240115",
			Status:        "In Progress",
			Period:        2,
			TimeRemaining: "5:32",
			HomeTeam: nba.Team{
				ID:           3,
				Name:         "Boston Celtics",
				Abbreviation: "BOS",
				Score:        58,
			},
			AwayTeam: nba.Team{
				ID:           4,
				Name:         "Miami Heat",
				Abbreviation: "MIA",
				Score:        52,
			},
			StartTime: time.Now(),
		},
	}

	filename := "test_nba_games.xlsx"
	defer os.Remove(filename) // Clean up after test

	// Test export
	err := exporter.ExportGames(games, filename)
	require.NoError(t, err)

	// Verify file was created
	_, err = os.Stat(filename)
	require.NoError(t, err)

	// Verify file content
	f, err := excelize.OpenFile(filename)
	require.NoError(t, err)
	defer f.Close()

	// Check sheet exists
	sheets := f.GetSheetList()
	assert.Contains(t, sheets, "NBA Games")

	// Check headers
	headerRow, err := f.GetRows("NBA Games")
	require.NoError(t, err)
	require.Len(t, headerRow, 3) // Headers + 2 data rows

	expectedHeaders := []string{"Game ID", "Date", "Away Team", "Away Score", "Home Team", "Home Score", "Status", "Period", "Time Remaining"}
	assert.Equal(t, expectedHeaders, headerRow[0])

	// Check first game data
	assert.Equal(t, "0022300123", headerRow[1][0])
	assert.Equal(t, "20240115", headerRow[1][1])
	assert.Contains(t, headerRow[1][2], "Golden State Warriors")
	assert.Contains(t, headerRow[1][2], "GSW")
	assert.Equal(t, "105", headerRow[1][3])
	assert.Contains(t, headerRow[1][4], "Los Angeles Lakers")
	assert.Contains(t, headerRow[1][4], "LAL")
	assert.Equal(t, "110", headerRow[1][5])
	assert.Equal(t, "Final", headerRow[1][6])
	assert.Equal(t, "4", headerRow[1][7])

	// Check second game data
	assert.Equal(t, "0022300124", headerRow[2][0])
	assert.Equal(t, "In Progress", headerRow[2][6])
	assert.Equal(t, "2", headerRow[2][7])
	assert.Equal(t, "5:32", headerRow[2][8])
}

func TestExportGames_EmptyGames(t *testing.T) {
	exporter := NewExcelExporter()
	filename := "empty_test.xlsx"
	defer os.Remove(filename)

	err := exporter.ExportGames([]nba.Game{}, filename)
	require.NoError(t, err)

	// Verify file was created with headers only
	f, err := excelize.OpenFile(filename)
	require.NoError(t, err)
	defer f.Close()

	rows, err := f.GetRows("NBA Games")
	require.NoError(t, err)
	assert.Len(t, rows, 1) // Only headers
}