# Variables
SRC_FILE := ./cmd/main.go
BUILD_DIR := ./tmp
APP_NAME := main

POSTGRES_CONTAINER := rss-postgres
REDIS_SERVER_CMD := redis-server
FORMAT_SCRIPT := scripts/format.sh
GOOSE_SCRIPT := scripts/goose.sh
NEXTJS_DIR := ./www

# Default target
.PHONY: ready
ready: start-postgres start-redis

# Build the project
.PHONY: build
build:
	@echo "Building the project..."
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_FILE)

# Run air for live reloading
.PHONY: run-air
run-air:
	@echo "Running air for live reloading..."
	air

# Start PostgreSQL container
.PHONY: start-postgres
start-postgres:
	@echo "Starting PostgreSQL container..."
	docker start $(POSTGRES_CONTAINER)

# Start Redis server
.PHONY: start-redis
start-redis:
	@echo "Starting Redis server..."
	$(REDIS_SERVER_CMD)

# Format the code
.PHONY: format
format:
	@echo "Formatting code..."
	$(FORMAT_SCRIPT)

# Goose migrations (up)
.PHONY: migrate-up
migrate-up:
	@echo "Running Goose migrations (up)..."
	$(GOOSE_SCRIPT) up

# Goose migrations (down)
.PHONY: migrate-down
migrate-down:
	@echo "Running Goose migrations (down)..."
	$(GOOSE_SCRIPT) down

# Start Next.js project (Bun)
.PHONY: start-nextjs
start-nextjs:
	@echo "Starting Next.js project..."
	cd $(NEXTJS_DIR) && bun run dev

# Clean build files
.PHONY: clean
clean:
	@echo "Cleaning build files..."
	rm -rf $(BUILD_DIR)

# Help
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  build           - Build the project"
	@echo "  start-postgres  - Start PostgreSQL container"
	@echo "  start-redis     - Start Redis server"
	@echo "  run-air         - Run air for live reloading"
	@echo "  format          - Format the code"
	@echo "  migrate-up      - Run Goose migrations (up)"
	@echo "  migrate-down    - Run Goose migrations (down)"
	@echo "  start-nextjs    - Start Next.js project (Bun)"
	@echo "  clean           - Clean build files"
