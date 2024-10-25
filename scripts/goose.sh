#!/bin/bash

# Load environment variables from the .env file
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
else
    echo ".env file not found."
    exit 1
fi

# Check if POSTGRES_URL is set
if [ -z "$POSTGRES_URL" ]; then
    echo "POSTGRES_URL not set in .env file."
    exit 1
fi

# Define the migration directory
MIGRATION_DIR="sql/schema"

# Check if the migration directory exists
if [ ! -d "$MIGRATION_DIR" ]; then
    echo "Migration directory '$MIGRATION_DIR' does not exist."
    exit 1
fi

# Check if an argument is provided
if [ $# -eq 0 ]; then
    echo "No arguments provided. Please use 'up' or 'down'."
    exit 1
fi

# Change to the migration directory
cd "$MIGRATION_DIR"

# Check if the argument is 'up' or 'down' and run goose accordingly
if [ "$1" = "up" ]; then
    echo "Running goose up"
    goose postgres "$POSTGRES_URL" up
elif [ "$1" = "down" ]; then
    echo "Running goose down"
    goose postgres "$POSTGRES_URL" down
else
    echo "Invalid argument. Please use 'up' or 'down'."
    exit 1
fi
