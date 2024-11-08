# Build stage using go's latest stable version
FROM golang:1.23.3-alpine AS go-builder

# Set the working directory inside the container
WORKDIR /src

# Install git for fetching dependencies
RUN apk add --no-cache git

# Set environment variables for Go
# Enable Go modules
ENV GO111MODULE=on  
# Disable CGO for static builds
ENV CGO_ENABLED=0
 # Set target OS
ENV GOOS=linux
# Set target architecture   
ENV GOARCH=amd64    

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
# This layer will be cached if the go.mod and go.sum files don't change
RUN go mod download

# Copy the entire source code
COPY . .

# Run tests
RUN go test -v ./...

# Build the application
# -ldflags="-w -s" reduces binary size by omitting debug information
# recommended only for release builds 
RUN go build -ldflags="-w -s" -o bin/main .

# create a minimal image for running the application
FROM alpine:3.18

# Set the working directory
WORKDIR /app

# Copy only the built binary from the build stage
COPY --from=go-builder /src/bin/main .

# Create a non-root user for running the application
RUN adduser -D appuser

# Switch to the non-root user
USER appuser

# container listens on port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]