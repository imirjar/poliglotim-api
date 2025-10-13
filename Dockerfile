# Stage 1: Builder
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
# CGO_ENABLED=0 disables CGO, creating a statically linked binary
# -o specifies the output file name
RUN GOOS=linux GOARCH=amd64 go build -o poliglotim-api cmd/main.go

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your application listens on (optional)
EXPOSE 8080

# Command to run the application
CMD ["./main"]