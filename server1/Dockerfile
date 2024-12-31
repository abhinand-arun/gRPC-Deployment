# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN GOOS=linux GOARCH=amd64 go build -o myapp .

# Stage 2: Create the final lightweight image
FROM alpine:latest

# Install necessary dependencies for running the Go binary (if needed)
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the port your application will listen on (if applicable)
EXPOSE 8080

# Set the command to run the application
CMD ["./myapp"]
