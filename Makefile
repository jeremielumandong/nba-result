# NBA Game Results Tracker Makefile

.PHONY: build test clean run install help

# Default target
all: test build

# Build the application
build:
	@echo "Building NBA Game Results Tracker..."
	go build -o bin/nba-tracker main.go
	@echo "Build complete: bin/nba-tracker"

# Run tests
test:
	@echo "Running tests..."
	go test ./tests/... -v
	go test ./internal/nba/... -v

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test ./tests/... ./internal/nba/... -v -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run the application with default settings
run:
	@echo "Running NBA Game Results Tracker..."
	go run main.go

# Run with custom date
run-date:
	@echo "Running NBA Game Results Tracker for specific date..."
	go run main.go -date 2024-01-15

# Run with date range
run-range:
	@echo "Running NBA Game Results Tracker for date range..."
	go run main.go -start-date 2024-01-15 -end-date 2024-01-17

# Install dependencies
install:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	rm -f *.json *.xlsx
	@echo "Clean complete"

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	golangci-lint run

# Run static analysis
vet:
	@echo "Running go vet..."
	go vet ./...

# Generate mock data for testing
mock-data:
	@echo "Running with mock data..."
	go run main.go -date 2024-01-15 -output mock_results.json -excel mock_report.xlsx

# Test date functionality
test-dates:
	@echo "Testing date functionality..."
	go run main.go -date 2024-01-15 -output test_single.json -excel test_single.xlsx
	go run main.go -start-date 2024-01-15 -end-date 2024-01-17 -output test_range.json -excel test_range.xlsx
	@echo "Test files generated: test_single.json, test_single.xlsx, test_range.json, test_range.xlsx"

# Development workflow
dev: fmt vet test build
	@echo "Development checks complete"

# Help target
help:
	@echo "Available targets:"
	@echo "  build         - Build the application"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  run           - Run the application with default settings"
	@echo "  run-date      - Run the application for a specific date"
	@echo "  run-range     - Run the application for a date range"
	@echo "  install       - Install dependencies"
	@echo "  clean         - Clean build artifacts"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code"
	@echo "  vet           - Run go vet"
	@echo "  mock-data     - Generate mock data files"
	@echo "  test-dates    - Test date functionality with sample outputs"
	@echo "  dev           - Run development workflow (fmt, vet, test, build)"
	@echo "  help          - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make run-date              # Run for 2024-01-15"
	@echo "  make run-range             # Run for 2024-01-15 to 2024-01-17"
	@echo "  make test-dates            # Test date functionality"