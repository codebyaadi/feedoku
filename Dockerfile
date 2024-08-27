# Use the official Golang image as the base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy the rest of the application code
COPY . .

# Build the Go app (specify the path to the main.go file)
RUN go build -v -o rss-app ./cmd

# Specify the command to run the app
CMD ["/app/rss-app"]
