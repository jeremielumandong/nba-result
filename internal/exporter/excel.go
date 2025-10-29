package exporter

import (
	"fmt"
	"strconv"

	"github.com/jeremielumandong/nba-result/internal/nba"
	"github.com/xuri/excelize/v2"
)

// ExcelExporter handles Excel file generation
type ExcelExporter struct{}

// NewExcelExporter creates a new Excel exporter
func NewExcelExporter() *ExcelExporter {
	return &ExcelExporter{}
}

// ExportGames exports NBA games to an Excel file
func (e *ExcelExporter) ExportGames(games []nba.Game, filename string) error {
	f := excelize.NewFile()
	defer f.Close()

	// Create sheet
	sheetName := "NBA Games"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("failed to create sheet: %w", err)
	}

	// Set headers
	headers := []string{
		"Game ID",
		"Date", 
		"Away Team",
		"Away Score",
		"Home Team",
		"Home Score",
		"Status",
		"Period",
		"Time Remaining",
	}

	// Write headers
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheetName, cell, header)
	}

	// Style headers
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#D3D3D3"},
			Pattern: 1,
		},
	})
	f.SetRowStyle(sheetName, 1, 1, headerStyle)

	// Write game data
	for i, game := range games {
		row := i + 2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), game.GameID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), game.Date)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), fmt.Sprintf("%s (%s)", game.AwayTeam.Name, game.AwayTeam.Abbreviation))
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), game.AwayTeam.Score)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), fmt.Sprintf("%s (%s)", game.HomeTeam.Name, game.HomeTeam.Abbreviation))
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), game.HomeTeam.Score)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), game.Status)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), game.Period)
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), game.TimeRemaining)
	}

	// Auto-adjust column widths
	for i := 0; i < len(headers); i++ {
		colName := string(rune('A' + i))
		f.SetColWidth(sheetName, colName, colName, 15)
	}

	// Set active sheet
	f.SetActiveSheet(index)

	// Save file
	if err := f.SaveAs(filename); err != nil {
		return fmt.Errorf("failed to save Excel file: %w", err)
	}

	return nil
}