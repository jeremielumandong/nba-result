# NBA Game Results Tracker Makefile

# Variables
APP_NAME=nba-tracker
GO_VERSION=1.21
BUILD_DIR=build
DIST_DIR=dist

# Default target
.PHONY: all
all: clean test build

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR) $(DIST_DIR) $(APP_NAME)
	rm -f *.json *.xlsx *.log

# Install dependencies
.PHONY: deps
deps:
	go mod tidy
	go mod download

# Run tests
.PHONY: test
test:
	go test -v ./tests/...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./tests/...
	go tool cover -html=coverage.out -o coverage.html

# Build the application
.PHONY: build
build: deps
	go build -o $(APP_NAME) main.go

# Build for multiple platforms
.PHONY: build-all
build-all: deps
	mkdir -p $(DIST_DIR)
	# Linux AMD64
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/$(APP_NAME)-linux-amd64 main.go
	# Linux ARM64
	GOOS=linux GOARCH=arm64 go build -o $(DIST_DIR)/$(APP_NAME)-linux-arm64 main.go
	# macOS AMD64
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/$(APP_NAME)-darwin-amd64 main.go
	# macOS ARM64 (Apple Silicon)
	GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/$(APP_NAME)-darwin-arm64 main.go
	# Windows AMD64
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/$(APP_NAME)-windows-amd64.exe main.go

# Install the application
.PHONY: install
install: build
	go install

# Run the application with today's games
.PHONY: run
run: build
	./$(APP_NAME)

# Run with specific date (usage: make run-date DATE=2024-01-15)
.PHONY: run-date
run-date: build
	./$(APP_NAME) -date=$(DATE)

# Run linter
.PHONY: lint
lint:
	golangci-lint run

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Check for security vulnerabilities
.PHONY: security
security:
	gosec ./...

# Generate documentation
.PHONY: docs
docs:
	godoc -http=:8080

# Development workflow
.PHONY: dev
dev: clean fmt lint test build

# Create release package
.PHONY: release
release: clean test build-all
	mkdir -p $(DIST_DIR)/release
	cp README.md $(DIST_DIR)/release/
	cp LICENSE $(DIST_DIR)/release/ 2>/dev/null || true
	cd $(DIST_DIR) && tar -czf release/$(APP_NAME)-linux-amd64.tar.gz $(APP_NAME)-linux-amd64 -C ../release README.md
	cd $(DIST_DIR) && tar -czf release/$(APP_NAME)-linux-arm64.tar.gz $(APP_NAME)-linux-arm64 -C ../release README.md
	cd $(DIST_DIR) && tar -czf release/$(APP_NAME)-darwin-amd64.tar.gz $(APP_NAME)-darwin-amd64 -C ../release README.md
	cd $(DIST_DIR) && tar -czf release/$(APP_NAME)-darwin-arm64.tar.gz $(APP_NAME)-darwin-arm64 -C ../release README.md
	cd $(DIST_DIR) && zip -r release/$(APP_NAME)-windows-amd64.zip $(APP_NAME)-windows-amd64.exe ../release/README.md

# Docker build
.PHONY: docker
docker:
	docker build -t $(APP_NAME) .

# Docker run
.PHONY: docker-run
docker-run:
	docker run --rm -v $(PWD)/output:/app/output $(APP_NAME)

# Show help
.PHONY: help
help:
	@echo "NBA Game Results Tracker - Available commands:"
	@echo ""
	@echo "  all          - Clean, test, and build (default)"
	@echo "  clean        - Remove build artifacts and output files"
	@echo "  deps         - Install/update dependencies"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  build        - Build the application"
	@echo "  build-all    - Build for multiple platforms"
	@echo "  install      - Install the application"
	@echo "  run          - Build and run with today's games"
	@echo "  run-date     - Build and run with specific date (make run-date DATE=2024-01-15)"
	@echo "  lint         - Run linter"
	@echo "  fmt          - Format code"
	@echo "  security     - Check for security vulnerabilities"
	@echo "  docs         - Start documentation server"
	@echo "  dev          - Full development workflow"
	@echo "  release      - Create release packages"
	@echo "  docker       - Build Docker image"
	@echo "  docker-run   - Run in Docker container"
	@echo "  help         - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make run                    # Fetch today's games"
	@echo "  make run-date DATE=2024-01-15  # Fetch games for specific date"
	@echo "  make test-coverage          # Run tests with coverage"
	@echo "  make build-all              # Build for all platforms"