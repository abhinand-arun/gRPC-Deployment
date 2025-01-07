package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kannan112/proto/ping"
	"google.golang.org/grpc"
)

const (
	// grpcServerAddress = "localhost:50051"
	defaultGrpcServerAddress = "localhost:50051"
)

func main() {

	r := gin.Default()

	// Define the /ping route
	r.GET("/ping", func(c *gin.Context) {

		// Create a gRPC client connection
		grpcServerAddress := os.Getenv("GRPC_SERVER_ADDRESS")
		if grpcServerAddress == "" {
			grpcServerAddress = defaultGrpcServerAddress
		}

		conn, err := grpc.Dial(grpcServerAddress, grpc.WithInsecure(), grpc.WithTimeout(10*time.Second))
		if err != nil {
			log.Fatalf("could not connect to gRPC server: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
			return
		}
		defer conn.Close()

		// Create a new PingService client
		client := ping.NewPingServiceClient(conn)

		// Create the PingRequest
		req := &ping.PingRequest{
			Message: "Hello from Gin HTTP server",
		}

		// Call the Ping method on the gRPC server
		rpcResponse, err := client.Ping(context.Background(), req)
		if err != nil {
			log.Printf("Error calling gRPC Ping: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response from gRPC server"})
			return
		}

		// Return the gRPC response as a JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": rpcResponse.GetReply(),
		})
	})

	r.Run(":8080")
}
