# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=gofmt
GOVET=$(GOCMD) vet
BINARY_NAME=nba-result
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: all build clean test coverage deps fmt vet run help

# Default target
all: test build

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./main.go

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Run tests
test:
	$(GOTEST) -v ./...

# Run tests with coverage
coverage:
	$(GOTEST) -cover ./...
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

# Download dependencies
deps:
	$(GOGET) -d ./...
	$(GOCMD) mod tidy

# Format code
fmt:
	$(GOFMT) -w .

# Vet code
vet:
	$(GOVET) ./...

# Run the application
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./main.go
	./$(BINARY_NAME)

# Run development server with auto-reload (requires air)
dev:
	air

# Build for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./main.go

# Docker build
docker-build:
	docker build -t $(BINARY_NAME):latest .

# Help
help:
	@echo "Available targets:"
	@echo "  build     - Build the application"
	@echo "  clean     - Clean build files"
	@echo "  test      - Run tests"
	@echo "  coverage  - Run tests with coverage report"
	@echo "  deps      - Download dependencies"
	@echo "  fmt       - Format code"
	@echo "  vet       - Vet code"
	@echo "  run       - Build and run the application"
	@echo "  dev       - Run with auto-reload (requires air)"
	@echo "  help      - Show this help message"
