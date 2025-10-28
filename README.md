# NBA Results Console App

A Go console application that fetches current NBA game results and generates both JSON and Excel reports.

## Features

- Fetch NBA game results for any specified date
- Generate JSON output with detailed game information
- Create formatted Excel reports with game statistics
- Support for current day results by default
- Command-line interface with customizable options
- Comprehensive error handling and logging

## Installation

### Prerequisites
- Go 1.21 or higher
- Internet connection to fetch NBA data

### Build from source

```bash
git clone https://github.com/jeremielumandong/nba-result.git
cd nba-result
go mod download
go build -o nba-results ./cmd/nba-results
```

## Usage

### Basic Usage

```bash
# Fetch today's NBA games
./nba-results

# Fetch games for a specific date
./nba-results -date 2024-01-15

# Specify custom output files
./nba-results -json games.json -excel games.xlsx

# Show help
./nba-results -help
```

### Command Line Options

- `-date`: Date to fetch games for (YYYY-MM-DD format, default: today)
- `-json`: Output JSON file path (default: nba_results.json)
- `-excel`: Output Excel file path (default: nba_results.xlsx)
- `-help`: Show help message

## Output Format

### JSON Output

The JSON output contains an array of game objects with the following structure:

```json
[
  {
    "game_id": "0022300567",
    "game_date": "2024-01-15",
    "game_time": "19:00",
    "home_team": {
      "name": "Lakers",
      "score": 108
    },
    "away_team": {
      "name": "Warriors",
      "score": 112
    },
    "status": "Final",
    "season_type": "Regular Season",
    "created_at": "2024-01-15T20:30:00Z"
  }
]
```

### Excel Output

The Excel report contains two sheets:

1. **NBA Games**: Detailed game information with the following columns:
   - Game Date
   - Game Time
   - Away Team
   - Away Score
   - Home Team
   - Home Score
   - Winner
   - Status
   - Game ID

2. **Summary**: Statistics including total games, finished games, and ongoing/scheduled games

## Architecture

The application is structured as follows:

```
├── main.go                 # Entry point and CLI handling
├── internal/
│   ├── nba/               # NBA API integration
│   │   ├── client.go      # HTTP client for NBA API
│   │   └── models.go      # Data models
│   └── report/            # Report generation
│       └── excel.go       # Excel report generator
├── tests/                 # Unit tests
└── README.md
```

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/nba
go test ./internal/report
```

## Error Handling

The application includes comprehensive error handling for:
- Network connectivity issues
- Invalid date formats
- API response parsing errors
- File system operations
- Excel report generation

## Dependencies

- `github.com/360EntSecGroup-Skylar/excelize/v2`: Excel file generation
- `github.com/stretchr/testify`: Testing framework

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## API Data Source

This application uses the NBA Stats API to fetch game data. Please note that the API usage should comply with NBA's terms of service.

## Support

For issues and questions, please create an issue in the GitHub repository.