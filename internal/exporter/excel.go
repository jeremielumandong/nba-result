package exporter

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jeremielumandong/nba-result/internal/nba"
	"github.com/xuri/excelize/v2"
)

// ExportToExcel exports games to an Excel file
func ExportToExcel(games []nba.Game, filename string) error {
	if len(games) == 0 {
		return fmt.Errorf("no games to export")
	}

	// Create new Excel file
	f := excelize.NewFile()

	// Create main sheet
	sheetName := "NBA Games"
	f.SetSheetName("Sheet1", sheetName)

	// Set headers
	headers := []string{
		"Game ID", "Date", "Home Team", "Home Score", 
		"Visitor Team", "Visitor Score", "Period", 
		"Clock", "Status", "Winner",
	}

	// Apply header styling
	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#366092"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#FFFFFF",
			Bold:  true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create header style: %w", err)
	}

	// Write headers
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
		f.SetCellStyle(sheetName, cell, cell, headerStyle)
	}

	// Create styles for different game statuses
	finalStyle, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#d4edda"},
			Pattern: 1,
		},
	})

	liveStyle, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#fff3cd"},
			Pattern: 1,
		},
	})

	scheduledStyle, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#f8f9fa"},
			Pattern: 1,
		},
	})

	// Write game data
	for i, game := range games {
		row := i + 2 // Starting from row 2 (after header)
		
		// Convert scores to integers for better display
		homeScore, _ := strconv.Atoi(game.HomeTeam.Score)
		visitorScore, _ := strconv.Atoi(game.VisitorTeam.Score)

		data := []interface{}{
			game.GameID,
			game.Date,
			game.HomeTeam.Tricode,
			homeScore,
			game.VisitorTeam.Tricode,
			visitorScore,
			game.Period,
			game.Clock,
			game.Status,
			game.Winner,
		}

		// Write data
		for j, value := range data {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			f.SetCellValue(sheetName, cell, value)
		}

		// Apply row styling based on game status
		var style int
		switch game.Status {
		case "Final":
			style = finalStyle
		case "Live":
			style = liveStyle
		case "Scheduled":
			style = scheduledStyle
		}

		// Apply style to entire row
		startCell, _ := excelize.CoordinatesToCellName(1, row)
		endCell, _ := excelize.CoordinatesToCellName(len(headers), row)
		f.SetCellStyle(sheetName, startCell, endCell, style)
	}

	// Auto-fit columns
	for i := range headers {
		colName, _ := excelize.ColumnNumberToName(i + 1)
		f.SetColWidth(sheetName, colName, colName, 15)
	}

	// Create summary sheet
	summarySheet := "Summary"
	f.NewSheet(summarySheet)

	// Count games by status
	var scheduled, live, final int
	for _, game := range games {
		switch game.Status {
		case "Scheduled":
			scheduled++
		case "Live":
			live++
		case "Final":
			final++
		}
	}

	// Write summary data
	summaryData := [][]interface{}{
		{"NBA Games Summary", ""},
		{"Generated", time.Now().Format("2006-01-02 15:04:05")},
		{"Date", games[0].Date},
		{"Total Games", len(games)},
		{"", ""},
		{"Status Breakdown", ""},
		{"Scheduled", scheduled},
		{"Live", live},
		{"Final", final},
	}

	for i, row := range summaryData {
		for j, value := range row {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+1)
			f.SetCellValue(summarySheet, cell, value)
		}
	}

	// Style summary sheet
	titleStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 16,
		},
	})
	f.SetCellStyle(summarySheet, "A1", "A1", titleStyle)

	// Set column widths for summary
	f.SetColWidth(summarySheet, "A", "A", 20)
	f.SetColWidth(summarySheet, "B", "B", 15)

	// Save file
	if err := f.SaveAs(filename); err != nil {
		return fmt.Errorf("failed to save Excel file: %w", err)
	}

	return nil
}