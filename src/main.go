package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sheva0914/go-practice/src/middleware"
	pb "github.com/sheva0914/go-practice/src/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func main() {
	log.Println("Server running...")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a server with middleware
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			middleware.AuthInterceptor(),
		),
	)
	// Register gRPC Server to the server
	pb.RegisterPracticeServer(s, &server{})

	// Listen and serve
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) SayHelloToWorld(ctx context.Context, in *pb.SayHelloToWorldRequest) (*pb.SayHelloToWorldResponse, error) {
	log.Printf("Received: %v\n", in.Name)
	message := fmt.Sprintf("Hello %v's world!\n", in.Name)
	log.Println(message)
	return &pb.SayHelloToWorldResponse{Message: message}, nil
}
