# Use the official Golang image as a builder
FROM golang:1.23 AS builder
# Set working directory inside the container
WORKDIR /app
# Copy the Go modules files
COPY go.mod go.sum ./
# Download dependencies
RUN go mod tidy
# Copy the source code
COPY . .
# Build the Go binary
RUN go build -o backend
# Use a minimal image for the final container
FROM alpine:latest
# Set working directory inside the container
WORKDIR /app
# Copy the built Go binary from the builder stage
COPY --from=builder /app/backend .
# Expose the API port
EXPOSE 8080
# Run the backend application
CMD ["./backend"]
