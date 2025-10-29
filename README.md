# NBA Game Results Tracker

A Go console application that fetches NBA game results for any date or date range and generates both JSON and Excel reports.

## Features

- **Date-specific queries**: Fetch NBA game results for any specific date
- **Date range queries**: Get games across multiple dates (up to 30 days)
- **Multiple output formats**: Generate JSON output and formatted Excel reports
- **Comprehensive validation**: Date format validation and business rule checks
- **Rich metadata**: Include summary statistics and generation metadata
- **Command-line interface**: Flexible options for different use cases
- **Mock data support**: Fallback to demonstration data when live APIs are unavailable

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
- `-start-date`: Start date for range query (YYYY-MM-DD)
- `-end-date`: End date for range query (YYYY-MM-DD)
- `-help`: Show help message

### Examples

**Single Date Queries:**
```bash
# Get today's games
go run main.go

# Get games for a specific date
go run main.go -date 2024-01-15

# Specify custom output files
go run main.go -date 2024-01-15 -output results.json -excel report.xlsx
```

**Date Range Queries:**
```bash
# Get games for a date range (3 days)
go run main.go -start-date 2024-01-15 -end-date 2024-01-17

# Date range with custom output files
go run main.go -start-date 2024-01-15 -end-date 2024-01-17 -output range_results.json -excel range_report.xlsx
```

**Help:**
```bash
go run main.go -help
```

## Output Formats

### JSON Format
The JSON output contains structured game results with metadata:
```json
{
  "date": "2024-01-15",
  "games": [
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
  ],
  "total_games": 1,
  "summary": {
    "scheduled": 0,
    "live": 0,
    "final": 1,
    "other": 0
  },
  "metadata": {
    "generated_at": "2024-01-16T10:30:00Z",
    "source": "NBA API",
    "version": "1.0"
  }
}
```

### Excel Report
The Excel report includes:
- Formatted table with all game details
- Winner determination for completed games
- Summary statistics (total games, games by status)
- Professional styling and auto-adjusted columns

## API and Architecture

### New Date Service
The application now includes a dedicated `DateService` that provides:
- Date validation and parsing
- Business rule enforcement (no future dates, no pre-NBA dates)
- Single date and date range query support
- Rich result metadata and summaries

### Key Components
- `DateService`: Handles date-based game queries with validation
- `Client`: NBA API interaction
- `ExcelReporter`: Excel report generation
- `GameResults`: Structured response format with metadata

## Validation Rules

- **Date format**: Must be YYYY-MM-DD
- **Future dates**: Not allowed
- **Historical limit**: No dates before 1946 (NBA founding year)
- **Range limit**: Maximum 30 days for range queries
- **Range logic**: End date must be after start date

## Project Structure

```
.
├── main.go                          # Main application with enhanced date functionality
├── go.mod                           # Go module definition
├── internal/
│   ├── nba/
│   │   ├── client.go                # NBA API client
│   │   ├── client_test.go           # Client tests
│   │   ├── date_service.go          # NEW: Date-based game queries
│   │   ├── date_service_test.go     # NEW: Date service tests
│   │   ├── date_types.go            # NEW: Date service types
│   │   ├── models.go                # Data models
│   │   └── types.go                 # Type definitions
│   ├── exporter/
│   │   ├── excel.go                 # Excel export functionality
│   │   ├── excel_test.go            # Excel export tests
│   │   └── json.go                  # JSON export functionality
│   └── report/
│       └── excel.go                 # Excel report generation
├── tests/
│   ├── exporter_test.go             # Exporter integration tests
│   ├── nba_test.go                  # NBA service integration tests
│   └── report_test.go               # Report generation tests
└── Makefile                         # Build and development tasks
```

## Development

### Build the application:
```bash
make build
```

### Run tests:
```bash
make test
```

### Run with coverage:
```bash
make test-coverage
```

### Format and lint:
```bash
make fmt
make vet
```

### Development workflow:
```bash
make dev
```

## Testing

The application includes comprehensive tests:
- Unit tests for date service functionality
- Integration tests for NBA API client
- Tests for Excel and JSON export
- Edge case and error condition testing

Run tests with:
```bash
go test ./... -v
```

## Error Handling

The application provides clear error messages for common issues:
- Invalid date formats
- Future or pre-NBA dates
- Network connectivity issues
- File system errors
- Date range validation errors

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- NBA for providing game data
- The Go community for excellent tooling and libraries
- Contributors and users of this project