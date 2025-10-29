# NBA Results API Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOVET=$(GOCMD) vet

# Binary names
BINARY_NAME=nba-results
BINARY_UNIX=$(BINARY_NAME)_unix

# Directories
CMD_DIR=./cmd/server
BIN_DIR=./bin

# Build the application
build:
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) -v $(CMD_DIR)

# Test all packages
test:
	$(GOTEST) -v ./...

# Test with coverage
test-coverage:
	$(GOTEST) -v -race -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -rf $(BIN_DIR)
	rm -f coverage.out coverage.html

# Run the application
run:
	$(GOCMD) run $(CMD_DIR)/main.go

# Install dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Format code
fmt:
	$(GOFMT) -s -w .

# Vet code for potential issues
vet:
	$(GOVET) ./...

# Lint code (requires golangci-lint to be installed)
lint:
	golangci-lint run

# Build for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BIN_DIR)/$(BINARY_UNIX) -v $(CMD_DIR)

# Build for multiple platforms
build-all: build build-linux

# Development setup
dev-setup:
	$(GOMOD) download
	$(GOMOD) tidy
	@echo "Installing development tools..."
	$(GOCMD) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Development setup complete!"

# Run all checks (format, vet, lint, test)
check: fmt vet lint test

# Create directories
init:
	mkdir -p $(BIN_DIR)

# Watch for changes and restart (requires air: go install github.com/cosmtrek/air@latest)
dev:
	air

# Help
help:
	@echo "Available targets:"
	@echo "  build        - Build the application"
	@echo "  test         - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  clean        - Clean build artifacts"
	@echo "  run          - Run the application"
	@echo "  deps         - Install dependencies"
	@echo "  fmt          - Format code"
	@echo "  vet          - Vet code for potential issues"
	@echo "  lint         - Lint code"
	@echo "  build-linux  - Build for Linux"
	@echo "  build-all    - Build for all platforms"
	@echo "  dev-setup    - Set up development environment"
	@echo "  check        - Run all checks (fmt, vet, lint, test)"
	@echo "  init         - Create necessary directories"
	@echo "  dev          - Watch for changes and restart (requires air)"
	@echo "  help         - Show this help message"

.PHONY: build test test-coverage clean run deps fmt vet lint build-linux build-all dev-setup check init dev help