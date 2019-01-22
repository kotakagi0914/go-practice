package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/sheva0914/go-practice/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func main() {
	fmt.Println("Server running...")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		os.Exit(1)
	}

	s := grpc.NewServer()
	pb.RegisterPracticeServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) SayHelloToWorld(ctx context.Context, in *pb.SayHelloToWorldRequest) (out *pb.SayHelloToWorldResponse, err error) {
	fmt.Printf("Received: %v\n", in.Name)
	message := fmt.Sprintf("Hello %v's world!\n", in.Name)
	fmt.Println(message)
	// out.Message = message
	return &pb.SayHelloToWorldResponse{Message: message}, nil
}
