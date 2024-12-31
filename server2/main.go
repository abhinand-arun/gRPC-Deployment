package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/kannan112/proto/ping"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement PingServiceServer
type server struct {
	ping.UnimplementedPingServiceServer
}

// Implement the Ping method of the PingService
func (s *server) Ping(ctx context.Context, req *ping.PingRequest) (*ping.PingResponse, error) {
	log.Printf("Received message: %s", req.GetMessage()) // Log the incoming message
	return &ping.PingResponse{
		Reply: "Hello from gRPC server!", // Respond with a message
	}, nil
}

func main() {
	// Listen for incoming gRPC requests on the specified port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the PingService with the gRPC server
	ping.RegisterPingServiceServer(s, &server{})

	// Start the gRPC server
	fmt.Println("Starting gRPC server on port", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
