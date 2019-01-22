package middleware

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

const (
	validTokenValue = "good"
)

/*
 * UnaryServerInterceptor returns interceptor for middlewares.
 */
func AuthInterceptor() grpc.UnaryServerInterceptor {
	return selfAuth
}

func selfAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println("will authrize with metadata")
	// Extract metadata from context.
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("ok: %v\n", ok)
	fmt.Printf("metadata: %v\n", md)
	// Check if metada exists.
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "metadata cannot be found")
	}
	// Extract token from metadata.
	token := md.Get("token")
	fmt.Printf("token: %v\n", token)
	// Check if token is valid
	if len(token) == 0 || token[0] != validTokenValue {
		fmt.Println("Given token was invalid")
		return nil, grpc.Errorf(codes.PermissionDenied, "Given token is invalid")
	}

	reply, err := handler(ctx, req)
	return reply, err
}
