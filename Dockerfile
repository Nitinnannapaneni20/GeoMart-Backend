# Use an official Golang image as the base
FROM golang:1.23

# Set working directory inside the container
WORKDIR /app

# Copy go modules files and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go binary
RUN go build -o backend ./main.go

# Expose the API port
EXPOSE 8080

# Run the backend application
CMD ["./backend"]
