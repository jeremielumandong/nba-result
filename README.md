# NBA Game Result Tracker

A console application that fetches current NBA game results and exports them to both JSON and Excel formats.

## Features

- Fetches real-time NBA game data from the official NBA API
- Outputs results as formatted JSON to console
- Exports game data to Excel spreadsheet
- Handles both completed and in-progress games
- Includes team information, scores, and game status

## Installation

1. Clone the repository:
```bash
git clone https://github.com/jeremielumandong/nba-result.git
cd nba-result
```

2. Install dependencies:
```bash
go mod download
```

## Usage

### Running the Application

```bash
go run cmd/main.go
```

The application will:
1. Fetch today's NBA games from the official API
2. Display the results as formatted JSON in the console
3. Export the data to an Excel file named `nba_games_YYYYMMDD.xlsx`

### Output Format

#### JSON Output (Console)
```json
[
  {
    "game_id": "0022300123",
    "date": "20240115",
    "home_team": {
      "id": 1,
      "name": "Los Angeles Lakers",
      "abbreviation": "LAL",
      "score": 110,
      "wins": 25,
      "losses": 18
    },
    "away_team": {
      "id": 2,
      "name": "Golden State Warriors", 
      "abbreviation": "GSW",
      "score": 105,
      "wins": 22,
      "losses": 21
    },
    "status": "Final",
    "period": 4,
    "time_remaining": "",
    "start_time": "2024-01-15T19:30:00Z"
  }
]
```

#### Excel Output
The Excel file contains a formatted spreadsheet with columns:
- Game ID
- Date
- Away Team
- Away Score  
- Home Team
- Home Score
- Status
- Period
- Time Remaining

## Architecture

The application is structured as follows:

```
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── nba/
│   │   ├── client.go        # NBA API client
│   │   └── client_test.go   # Client tests
│   └── exporter/
│       ├── excel.go         # Excel export functionality
│       └── excel_test.go    # Excel export tests
├── go.mod                   # Go module dependencies
└── README.md               # Documentation
```

### Components

#### NBA Client (`internal/nba/client.go`)
- Handles HTTP requests to the NBA Stats API
- Parses API responses into structured data
- Manages authentication headers and request formatting
- Converts raw API data to Go structs

#### Excel Exporter (`internal/exporter/excel.go`)
- Creates formatted Excel spreadsheets
- Applies styling to headers and data
- Auto-adjusts column widths for readability
- Handles multiple games in a single export

## Dependencies

- `github.com/xuri/excelize/v2` - Excel file generation
- `github.com/stretchr/testify` - Testing framework

## Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run specific package tests:
```bash
go test ./internal/nba
go test ./internal/exporter
```

## API Information

This application uses the official NBA Stats API:
- Base URL: `https://stats.nba.com/stats/scoreboardV2`
- Requires specific headers to avoid rate limiting
- Returns comprehensive game and team data
- Updates in real-time during games

## Error Handling

The application includes robust error handling for:
- Network connectivity issues
- API rate limiting or unavailability
- Invalid API responses
- File system operations
- Data parsing errors

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/nba-game-tracker`)
3. Make your changes
4. Add tests for new functionality
5. Run the test suite
6. Commit your changes (`git commit -am 'Add new feature'`)
7. Push to the branch (`git push origin feature/nba-game-tracker`)
8. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
