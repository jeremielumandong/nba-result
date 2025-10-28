# NBA Game Results Tracker

A Go console application that fetches current NBA game results for any given day and exports them to both JSON and Excel formats.

## Features

- Fetch NBA game results for any date
- Export results to JSON format with metadata
- Generate formatted Excel reports with styling
- Support for live games, scheduled games, and final results
- Command-line interface with flexible options
- Comprehensive error handling and logging

## Installation

1. Clone the repository:
```bash
git clone https://github.com/jeremielumandong/nba-result.git
cd nba-result
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the application:
```bash
go build -o nba-tracker main.go
```

## Usage

### Basic Usage

Fetch today's NBA games:
```bash
./nba-tracker
```

### Options

```bash
./nba-tracker [options]
```

**Available Options:**
- `-date`: Specific date to fetch games for (YYYY-MM-DD format, defaults to today)
- `-json`: Output JSON file path (default: "nba_games.json")
- `-excel`: Output Excel file path (default: "nba_games.xlsx")
- `-help`: Show help message

### Examples

1. Fetch games for a specific date:
```bash
./nba-tracker -date=2024-01-15
```

2. Custom output file names:
```bash
./nba-tracker -json=results.json -excel=report.xlsx
```

3. Fetch games for yesterday with custom files:
```bash
./nba-tracker -date=2024-01-14 -json=yesterday_games.json -excel=yesterday_report.xlsx
```

## Output Formats

### JSON Output

The JSON output includes:
- Export metadata (generation time, total games, date)
- Complete game details (teams, scores, status, etc.)
- Summary statistics (games by status)

### Excel Output

The Excel file contains:
- **NBA Games Sheet**: Detailed game information with color-coded status
  - Green: Final games
  - Yellow: Live games
  - Gray: Scheduled games
- **Summary Sheet**: Overview statistics and metadata

## Game Status Types

- **Scheduled**: Game hasn't started yet
- **Live**: Game is currently in progress
- **Final**: Game has ended

## API Data Source

This application uses the unofficial NBA.com API endpoints to fetch game data. The API provides real-time information about:
- Game schedules
- Live scores and game clock
- Final results
- Team information

## Error Handling

The application includes comprehensive error handling for:
- Network connectivity issues
- Invalid date formats
- API response errors
- File creation/writing errors
- Missing or malformed data

## Dependencies

- [excelize/v2](https://github.com/xuri/excelize): Excel file generation
- Standard Go libraries for HTTP requests and JSON handling

## Development

### Running Tests

```bash
go test ./...
```

### Running with Verbose Output

```bash
go run main.go -date=2024-01-15
```

### Project Structure

```
.
├── main.go                 # Application entry point
├── internal/
│   ├── nba/
│   │   ├── client.go       # NBA API client
│   │   └── models.go       # Data structures
│   └── exporter/
│       ├── json.go         # JSON export functionality
│       └── excel.go        # Excel export functionality
├── tests/
│   ├── nba_test.go         # NBA client tests
│   └── exporter_test.go    # Exporter tests
├── go.mod                  # Go module file
└── README.md
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

MIT License - see LICENSE file for details

## Troubleshooting

### Common Issues

1. **"No games found"**: Check if the date is valid and if there were NBA games scheduled for that date
2. **"Failed to fetch games"**: Verify internet connectivity and try again
3. **"Permission denied"**: Ensure you have write permissions in the output directory

### Rate Limiting

The NBA API may have rate limiting. If you encounter issues, wait a few minutes before making another request.

### Date Format

Always use YYYY-MM-DD format for dates. Examples:
- 2024-01-15 ✓
- 01-15-2024 ✗
- 2024/01/15 ✗