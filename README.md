# NBA Results API

A clean and well-structured Go application for fetching and managing NBA game results.

## Features

- Fetch NBA game results
- RESTful API endpoints
- Clean architecture with separation of concerns
- Comprehensive error handling
- Unit tests with good coverage

## Project Structure

```
├── cmd/
│   └── server/          # Application entry points
├── internal/
│   ├── config/          # Configuration management
│   ├── handlers/        # HTTP handlers
│   ├── models/          # Data models
│   ├── services/        # Business logic
│   └── repository/      # Data access layer
├── pkg/
│   └── utils/           # Shared utilities
├── tests/               # Integration tests
└── docs/                # Documentation
```

## Getting Started

1. Clone the repository
2. Run `go mod tidy` to install dependencies
3. Run `go run cmd/server/main.go` to start the server
4. Access the API at `http://localhost:8080`

## API Endpoints

- `GET /api/games` - Get all games
- `GET /api/games/{id}` - Get game by ID
- `GET /api/teams/{team}/games` - Get games for a specific team

## Testing

Run tests with:
```bash
go test ./...
```

## Contributing

Please ensure code follows Go best practices and includes appropriate tests.