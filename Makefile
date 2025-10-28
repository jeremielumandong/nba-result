# NBA Result Tracker Makefile

.PHONY: build run test clean lint fmt vet install-deps

# Build the application
build:
	@echo "Building NBA Result Tracker..."
	go build -o bin/nba-result cmd/main.go

# Run the application
run:
	@echo "Running NBA Result Tracker..."
	go run cmd/main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf bin/
	rm -f *.xlsx
	rm -f coverage.out coverage.html

# Lint the code
lint:
	@echo "Running linter..."
	golangci-lint run

# Format the code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Vet the code
vet:
	@echo "Vetting code..."
	go vet ./...

# Install dependencies
install-deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Development workflow
dev: fmt vet test

# CI workflow
ci: fmt vet test-coverage

# Help
help:
	@echo "Available targets:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  clean        - Clean build artifacts"
	@echo "  lint         - Run linter"
	@echo "  fmt          - Format code"
	@echo "  vet          - Vet code"
	@echo "  install-deps - Install dependencies"
	@echo "  dev          - Run development workflow (fmt, vet, test)"
	@echo "  ci           - Run CI workflow (fmt, vet, test with coverage)"
	@echo "  help         - Show this help message"

default: help