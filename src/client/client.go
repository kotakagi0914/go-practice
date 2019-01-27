package main

import (
	"context"
	"log"
	"os"

	pb "github.com/sheva0914/go-practice/src/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address     = "localhost:50051"
	defaultName = "Kotaro"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewPracticeClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("token", "good"))

	res, err := client.SayHelloToWorld(ctx, &pb.SayHelloToWorldRequest{Name: name})
	if err != nil {
		log.Fatalf("Failed to request: %v", err)
	}
	log.Printf("Message from server: %s", res.Message)
}
