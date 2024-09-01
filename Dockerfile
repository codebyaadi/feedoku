# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache build-base

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy the rest of the application code
COPY . .

# Build the Go app (specify the path to the main.go file)
RUN go build -v -o rss-app ./cmd

# Stage 2: Create a minimal image to run the Go application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/rss-app /app/rss-app

# Specify the command to run the app
CMD ["/app/rss-app"]
