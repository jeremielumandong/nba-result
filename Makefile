# Makefile for NBA Results Console App

# Variables
APP_NAME := nba-results
BUILD_DIR := bin
MAIN_PATH := ./main.go
GO_VERSION := 1.21

# Default target
.DEFAULT_GOAL := help

# Build the application
.PHONY: build
build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(APP_NAME)"

# Build for multiple platforms
.PHONY: build-all
build-all: ## Build for multiple platforms
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME)-darwin-arm64 $(MAIN_PATH)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe $(MAIN_PATH)
	@echo "Multi-platform build complete"

# Run the application
.PHONY: run
run: ## Run the application
	go run $(MAIN_PATH)

# Run with specific date
.PHONY: run-date
run-date: ## Run the application with a specific date (make run-date DATE=2024-01-15)
	go run $(MAIN_PATH) -date $(DATE)

# Run tests
.PHONY: test
test: ## Run all tests
	go test ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage: ## Run tests with coverage report
	go test -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run tests with verbose output
.PHONY: test-verbose
test-verbose: ## Run tests with verbose output
	go test -v ./...

# Clean build artifacts
.PHONY: clean
clean: ## Clean build artifacts and generated files
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	rm -f *.json *.xlsx
	@echo "Clean complete"

# Install dependencies
.PHONY: deps
deps: ## Download and install dependencies
	go mod download
	go mod tidy

# Format code
.PHONY: fmt
fmt: ## Format Go code
	go fmt ./...

# Lint code
.PHONY: lint
lint: ## Run linter (requires golangci-lint)
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Vet code
.PHONY: vet
vet: ## Run go vet
	go vet ./...

# Security scan
.PHONY: security
security: ## Run security scan (requires gosec)
	@if command -v gosec >/dev/null 2>&1; then \
		gosec ./...; \
	else \
		echo "gosec not installed. Install it with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest"; \
	fi

# Development setup
.PHONY: dev-setup
dev-setup: ## Set up development environment
	@echo "Setting up development environment..."
	go mod download
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
	@echo "Development setup complete"

# Check Go version
.PHONY: check-go-version
check-go-version: ## Check Go version
	@go version
	@echo "Required Go version: $(GO_VERSION)"

# Create example output
.PHONY: example
example: build ## Create example output files
	@echo "Creating example output..."
	./$(BUILD_DIR)/$(APP_NAME) -json example_results.json -excel example_results.xlsx
	@echo "Example files created"

# Show help
.PHONY: help
help: ## Show this help message
	@echo "NBA Results Console App - Makefile Commands"
	@echo ""
	@awk 'BEGIN {FS = ":.*##"}; /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)