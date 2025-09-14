# Go Validator API Makefile

.PHONY: help dev build run clean swagger test

# Default target
help:
	@echo "Available commands:"
	@echo "  dev       - Run in development mode with hot reload"
	@echo "  build     - Build the application"
	@echo "  run       - Run the application"
	@echo "  clean     - Clean up Docker resources"
	@echo "  swagger   - Generate Swagger documentation"
	@echo "  test      - Run tests"
	@echo "  deps      - Download all dependencies"

# Run in development mode with hot reload
dev:
	docker compose up --build

# Build the application
build:
	go build -o bin/main .

# Run the application
run:
	go run .

# Clean up Docker resources
clean:
	docker compose down -v
	docker system prune -f

# Generate Swagger documentation
swagger:
	~/go/bin/swag init
	@echo "Swagger documentation generated at /swagger/index.html"

# Run tests
test:
	go test ./...

# Install dependencies
deps:
	go mod tidy
	go mod download
