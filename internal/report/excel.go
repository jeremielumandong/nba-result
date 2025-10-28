package report

import (
	"fmt"
	"strconv"

	"github.com/jeremielumandong/nba-result/internal/nba"
	"github.com/xuri/excelize/v2"
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

// GenerateReport generates an Excel report from NBA games data
func (r *ExcelReporter) GenerateReport(games []nba.Game, filename string) error {
	sheetName := "NBA Games"
	index, err := r.file.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("creating sheet: %w", err)
	}

	// Set the sheet as active
	r.file.SetActiveSheet(index)

	// Create headers
	headers := []string{
		"Game ID", "Date", "Time", "Away Team", "Away Score",
		"Home Team", "Home Score", "Status", "Quarter", "Time Left", "Winner",
	}

	// Set headers
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string(rune('A'+i)))
		if err := r.file.SetCellValue(sheetName, cell, header); err != nil {
			return fmt.Errorf("setting header %s: %w", header, err)
		}
	}

	// Style headers
	if err := r.styleHeaders(sheetName, len(headers)); err != nil {
		return fmt.Errorf("styling headers: %w", err)
	}

	// Add data
	for i, game := range games {
		row := i + 2 // Start from row 2 (after headers)
		
		data := []interface{}{
			game.GameID,
			game.Date,
			game.Time,
			game.AwayTeam.Name,
			game.AwayTeam.Score,
			game.HomeTeam.Name,
			game.HomeTeam.Score,
			game.Status,
			game.Quarter,
			game.TimeLeft,
			r.determineWinner(game),
		}

		for j, value := range data {
			cell := fmt.Sprintf("%s%d", string(rune('A'+j)), row)
			if err := r.file.SetCellValue(sheetName, cell, value); err != nil {
				return fmt.Errorf("setting cell %s: %w", cell, err)
			}
		}
	}

	// Auto-adjust column widths
	if err := r.autoAdjustColumns(sheetName, len(headers)); err != nil {
		return fmt.Errorf("auto-adjusting columns: %w", err)
	}

	// Add summary statistics
	if err := r.addSummary(sheetName, games, len(games)+3); err != nil {
		return fmt.Errorf("adding summary: %w", err)
	}

	// Delete default sheet
	if err := r.file.DeleteSheet("Sheet1"); err != nil {
		return fmt.Errorf("deleting default sheet: %w", err)
	}

	// Save the file
	return r.file.SaveAs(filename)
}

// styleHeaders applies styling to the header row
func (r *ExcelReporter) styleHeaders(sheetName string, numCols int) error {
	style, err := r.file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#4472C4"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#FFFFFF",
			Bold:  true,
		},
	})
	if err != nil {
		return err
	}

	// Apply style to header range
	headerRange := fmt.Sprintf("A1:%s1", string(rune('A'+numCols-1)))
	return r.file.SetCellStyle(sheetName, "A1", string(rune('A'+numCols-1))+"1", style)
}

// autoAdjustColumns adjusts column widths automatically
func (r *ExcelReporter) autoAdjustColumns(sheetName string, numCols int) error {
	colWidths := map[string]float64{
		"A": 12, // Game ID
		"B": 12, // Date
		"C": 8,  // Time
		"D": 20, // Away Team
		"E": 12, // Away Score
		"F": 20, // Home Team
		"G": 12, // Home Score
		"H": 12, // Status
		"I": 10, // Quarter
		"J": 12, // Time Left
		"K": 20, // Winner
	}

	for i := 0; i < numCols; i++ {
		col := string(rune('A' + i))
		if width, ok := colWidths[col]; ok {
			if err := r.file.SetColWidth(sheetName, col, col, width); err != nil {
				return err
			}
		}
	}
	return nil
}

// addSummary adds summary statistics to the report
func (r *ExcelReporter) addSummary(sheetName string, games []nba.Game, startRow int) error {
	// Count games by status
	statusCount := make(map[string]int)
	for _, game := range games {
		statusCount[game.Status]++
	}

	// Add summary title
	summaryTitleCell := fmt.Sprintf("A%d", startRow)
	if err := r.file.SetCellValue(sheetName, summaryTitleCell, "SUMMARY"); err != nil {
		return err
	}

	// Style summary title
	summaryStyle, err := r.file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 14,
		},
	})
	if err != nil {
		return err
	}
	if err := r.file.SetCellStyle(sheetName, summaryTitleCell, summaryTitleCell, summaryStyle); err != nil {
		return err
	}

	// Add total games
	totalCell := fmt.Sprintf("A%d", startRow+2)
	if err := r.file.SetCellValue(sheetName, totalCell, fmt.Sprintf("Total Games: %d", len(games))); err != nil {
		return err
	}

	// Add games by status
	row := startRow + 3
	for status, count := range statusCount {
		statusCell := fmt.Sprintf("A%d", row)
		if err := r.file.SetCellValue(sheetName, statusCell, fmt.Sprintf("%s Games: %d", status, count)); err != nil {
			return err
		}
		row++
	}

	return nil
}

// determineWinner determines the winner of a game
func (r *ExcelReporter) determineWinner(game nba.Game) string {
	if game.Status != "Final" {
		return "TBD"
	}

	if game.HomeTeam.Score > game.AwayTeam.Score {
		return game.HomeTeam.Name
	} else if game.AwayTeam.Score > game.HomeTeam.Score {
		return game.AwayTeam.Name
	}
	return "Tie"
}