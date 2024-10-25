#!/bin/bash

# Define environment variables
POSTGRES_USER="myuser"
POSTGRES_PASSWORD="mypassword"
POSTGRES_DB="rss-feed-agg"
CONTAINER_NAME="rss-postgres"
IMAGE="postgres:13-alpine"
PORT="5432"

# Remove any existing container with the same name
if docker ps -a --format '{{.Names}}' | grep -q "^${CONTAINER_NAME}$"; then
    echo "Removing existing container with name ${CONTAINER_NAME}..."
    docker rm -f ${CONTAINER_NAME}
fi

# Run the PostgreSQL container
echo "Starting PostgreSQL container..."
docker run -d \
  --name ${CONTAINER_NAME} \
  -e POSTGRES_USER=${POSTGRES_USER} \
  -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
  -e POSTGRES_DB=${POSTGRES_DB} \
  -p ${PORT}:5432 \
  ${IMAGE}

echo "PostgreSQL container ${CONTAINER_NAME} started successfully."
