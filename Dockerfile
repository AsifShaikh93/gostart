# Step 1: Use the official Golang image as builder
FROM golang:1.24.2 AS builder

# Set working directory inside container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Step 2: Create a small final image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the pre-built binary from builder
COPY --from=builder /app/main .

# Expose port (same as your Go server)
EXPOSE 8083

# Command to run the app
CMD ["./main"]
