# NBA Result API

A clean, well-structured Go REST API for NBA game results and team information.

## Features

- Get NBA game results
- Retrieve team information
- RESTful API design
- Comprehensive test coverage
- Clean code architecture

## Project Structure

```
nba-result/
├── main.go                     # Application entry point
├── go.mod                      # Go module definition
├── internal/                   # Internal packages
│   ├── api/                   # API routing
│   │   └── router.go
│   ├── config/                # Configuration management
│   │   └── config.go
│   ├── handlers/              # HTTP handlers
│   │   ├── health.go
│   │   ├── games.go
│   │   ├── teams.go
│   │   └── *_test.go         # Handler tests
│   ├── models/                # Data models
│   │   ├── game.go
│   │   ├── team.go
│   │   └── *_test.go         # Model tests
│   └── services/              # Business logic
│       ├── game_service.go
│       ├── team_service.go
│       └── *_test.go         # Service tests
└── README.md
```

## Architecture

This project follows clean architecture principles:

- **Handlers**: HTTP request/response handling
- **Services**: Business logic layer
- **Models**: Domain entities and data structures
- **Config**: Application configuration

## API Endpoints

### Health Check
- `GET /api/v1/health` - System health check

### Games
- `GET /api/v1/games` - Get all games
- `GET /api/v1/games/{id}` - Get game by ID

### Teams
- `GET /api/v1/teams` - Get all teams
- `GET /api/v1/teams/{id}` - Get team by ID

## Getting Started

### Prerequisites
- Go 1.21 or higher

### Installation

1. Clone the repository:
```bash
git clone https://github.com/jeremielumandong/nba-result.git
cd nba-result
```

2. Download dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The server will start on port 8080 by default.

### Configuration

Environment variables:
- `PORT`: Server port (default: 8080)
- `NBA_API_URL`: External NBA API URL (default: https://api.nba.com)

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

### Example Requests

```bash
# Health check
curl http://localhost:8080/api/v1/health

# Get all games
curl http://localhost:8080/api/v1/games

# Get specific game
curl http://localhost:8080/api/v1/games/1

# Get all teams
curl http://localhost:8080/api/v1/teams

# Get specific team
curl http://localhost:8080/api/v1/teams/1
```

## Code Quality

- **Clean Code**: Well-structured, readable, and maintainable
- **Testing**: Comprehensive unit tests for all components
- **Error Handling**: Proper error handling throughout
- **Documentation**: Clear comments and documentation
- **Go Best Practices**: Follows Go conventions and idioms

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add/update tests
5. Run tests to ensure they pass
6. Submit a pull request

## License

This project is licensed under the MIT License.
