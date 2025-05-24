FROM golang:1.20-alpine

# Create app directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod .
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["./main"]