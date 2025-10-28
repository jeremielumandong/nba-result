# NBA Game Results Tracker

A Go console application that fetches current NBA game results and generates both JSON and Excel reports.

## Features

- Fetch NBA game results for any date
- Generate JSON output with game details
- Create formatted Excel reports with statistics
- Command-line interface with flexible options
- Mock data support for demonstration (since free NBA APIs are limited)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/jeremielumandong/nba-result.git
cd nba-result
```

2. Initialize Go modules:
```bash
go mod tidy
```

## Usage

### Basic Usage

Run with default settings (today's games):
```bash
go run main.go
```

### Command Line Options

- `-output`: Specify JSON output file (default: `nba_results.json`)
- `-excel`: Specify Excel output file (default: `nba_results.xlsx`)
- `-date`: Specify date in YYYY-MM-DD format (default: today)
- `-help`: Show help message

### Examples

```bash
# Get today's games
go run main.go

# Get games for a specific date
go run main.go -date 2024-01-15

# Specify custom output files
go run main.go -output results.json -excel report.xlsx

# Get help
go run main.go -help
```

## Output

### JSON Format
The JSON output contains an array of game objects with the following structure:
```json
[
  {
    "game_id": "001",
    "date": "2024-01-15",
    "time": "20:00",
    "home_team": {
      "name": "Los Angeles Lakers",
      "code": "LAL",
      "score": 112
    },
    "away_team": {
      "name": "Boston Celtics",
      "code": "BOS",
      "score": 108
    },
    "status": "Final",
    "quarter": 4,
    "time_left": "0:00"
  }
]
```

### Excel Report
The Excel report includes:
- Formatted table with all game details
- Winner determination for completed games
- Summary statistics (total games, games by status)
- Professional styling and auto-adjusted columns

## Project Structure

```
.
├── main.go                 # Main application entry point
├── go.mod                  # Go module definition
├── internal/
│   ├── nba/
│   │   ├── client.go       # NBA API client
│   │   └── types.go        # Data structures
│   └── report/
│       └── excel.go        # Excel report generation
└── tests/
    ├── nba_test.go         # NBA client tests
    └── report_test.go      # Report generation tests
```

## Dependencies

- `github.com/xuri/excelize/v2`: Excel file generation
- `github.com/stretchr/testify`: Testing utilities

## Testing

Run tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

## API Integration

Currently, the application uses mock data for demonstration purposes. The NBA provides various APIs, but most require authentication or have rate limits. The mock data structure follows the expected format and can be easily replaced with real API integration.

To integrate with a real NBA API:
1. Update the `client.go` file with the appropriate API endpoints
2. Modify the parsing logic in `parseGamesFromAPI` method
3. Add proper error handling for API responses

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Make your changes and add tests
4. Run tests: `go test ./...`
5. Commit your changes: `git commit -am 'Add your feature'`
6. Push to the branch: `git push origin feature/your-feature`
7. Create a Pull Request

## License

This project is licensed under the MIT License.
