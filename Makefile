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

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test ./tests/... -v -cover -coverprofile=coverage.out
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
	go run main.go -output mock_results.json -excel mock_report.xlsx

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
	@echo "  install       - Install dependencies"
	@echo "  clean         - Clean build artifacts"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code (requires golangci-lint)"
	@echo "  vet           - Run go vet"
	@echo "  mock-data     - Generate sample output files"
	@echo "  dev           - Run development checks (fmt, vet, test, build)"
	@echo "  help          - Show this help message"
