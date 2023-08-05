# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy package files
COPY go.mod go.sum ./

# Installs Go dependencies
RUN go mod download

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o /application

# This container exposes port 8080 to the outside world
EXPOSE 8000

# Run the binary program produced by `go install`
CMD ["/application"]