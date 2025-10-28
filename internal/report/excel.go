// Package report provides functionality to generate reports from NBA game data
package report

import (
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/jeremielumandong/nba-result/internal/nba"
)

// ExcelReporter handles Excel report generation
type ExcelReporter struct {
	file *excelize.File
}

// NewExcelReporter creates a new Excel reporter
func NewExcelReporter() *ExcelReporter {
	return &ExcelReporter{
		file: excelize.NewFile(),
	}
}

// GenerateReport creates an Excel report from NBA games data
func (r *ExcelReporter) GenerateReport(games []nba.Game, filename string) error {
	sheetName := "NBA Games"
	r.file.SetSheetName("Sheet1", sheetName)

	// Set headers
	headers := []string{
		"Game Date", "Game Time", "Away Team", "Away Score", 
		"Home Team", "Home Score", "Winner", "Status", "Game ID",
	}

	// Apply header styling
	headerStyle, err := r.createHeaderStyle()
	if err != nil {
		return fmt.Errorf("failed to create header style: %w", err)
	}

	// Set headers
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string(rune('A'+i)))
		r.file.SetCellValue(sheetName, cell, header)
		r.file.SetCellStyle(sheetName, cell, cell, headerStyle)
	}

	// Add data rows
	for i, game := range games {
		row := i + 2 // Start from row 2 (after headers)
		
		// Set cell values
		r.file.SetCellValue(sheetName, fmt.Sprintf("A%d", row), game.GameDate)
		r.file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), game.GameTime)
		r.file.SetCellValue(sheetName, fmt.Sprintf("C%d", row), game.AwayTeam.Name)
		r.file.SetCellValue(sheetName, fmt.Sprintf("D%d", row), game.AwayTeam.Score)
		r.file.SetCellValue(sheetName, fmt.Sprintf("E%d", row), game.HomeTeam.Name)
		r.file.SetCellValue(sheetName, fmt.Sprintf("F%d", row), game.HomeTeam.Score)
		r.file.SetCellValue(sheetName, fmt.Sprintf("G%d", row), game.Winner())
		r.file.SetCellValue(sheetName, fmt.Sprintf("H%d", row), game.Status)
		r.file.SetCellValue(sheetName, fmt.Sprintf("I%d", row), game.GameID)

		// Apply row styling
		if err := r.applyRowStyling(sheetName, row, game.IsFinished()); err != nil {
			return fmt.Errorf("failed to apply row styling: %w", err)
		}
	}

	// Auto-fit columns
	for i := 0; i < len(headers); i++ {
		col := string(rune('A' + i))
		r.file.SetColWidth(sheetName, col, col, 15)
	}

	// Add summary sheet
	if err := r.addSummarySheet(games); err != nil {
		return fmt.Errorf("failed to add summary sheet: %w", err)
	}

	// Save file
	if err := r.file.SaveAs(filename); err != nil {
		return fmt.Errorf("failed to save Excel file: %w", err)
	}

	return nil
}

// createHeaderStyle creates styling for header row
func (r *ExcelReporter) createHeaderStyle() (int, error) {
	return r.file.NewStyle(&excelize.Style{
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
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
	})
}

// applyRowStyling applies styling to data rows
func (r *ExcelReporter) applyRowStyling(sheetName string, row int, isFinished bool) error {
	var fillColor string
	if isFinished {
		fillColor = "#E8F5E8" // Light green for finished games
	} else {
		fillColor = "#FFF2CC" // Light yellow for ongoing/scheduled games
	}

	style, err := r.file.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{fillColor},
			Pattern: 1,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#CCCCCC", Style: 1},
			{Type: "top", Color: "#CCCCCC", Style: 1},
			{Type: "bottom", Color: "#CCCCCC", Style: 1},
			{Type: "right", Color: "#CCCCCC", Style: 1},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	if err != nil {
		return err
	}

	// Apply style to the entire row
	for col := 'A'; col <= 'I'; col++ {
		cell := fmt.Sprintf("%c%d", col, row)
		r.file.SetCellStyle(sheetName, cell, cell, style)
	}

	return nil
}

// addSummarySheet creates a summary sheet with statistics
func (r *ExcelReporter) addSummarySheet(games []nba.Game) error {
	summarySheet := "Summary"
	r.file.NewSheet(summarySheet)

	// Calculate statistics
	totalGames := len(games)
	finishedGames := 0
	for _, game := range games {
		if game.IsFinished() {
			finishedGames++
		}
	}

	// Add summary data
	r.file.SetCellValue(summarySheet, "A1", "NBA Games Summary")
	r.file.SetCellValue(summarySheet, "A2", "Generated:")
	r.file.SetCellValue(summarySheet, "B2", time.Now().Format("2006-01-02 15:04:05"))
	r.file.SetCellValue(summarySheet, "A3", "Total Games:")
	r.file.SetCellValue(summarySheet, "B3", totalGames)
	r.file.SetCellValue(summarySheet, "A4", "Finished Games:")
	r.file.SetCellValue(summarySheet, "B4", finishedGames)
	r.file.SetCellValue(summarySheet, "A5", "Ongoing/Scheduled:")
	r.file.SetCellValue(summarySheet, "B5", totalGames-finishedGames)

	// Style the summary
	titleStyle, _ := r.file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 16,
		},
	})
	r.file.SetCellStyle(summarySheet, "A1", "A1", titleStyle)

	return nil
}