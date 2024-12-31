# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go binary
RUN GOOS=linux GOARCH=amd64 go build -o server .

# Stage 2: Create a lightweight image for running the application
FROM alpine:latest

# Set the working directo    ▪ Using image gcr.io/k8s-minikube/storage-provisioner:v5ry
WORKDIR /app


# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Expose the gRPC server port
EXPOSE 50051

# Command to run the application
CMD ["./server"]
