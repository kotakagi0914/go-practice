package middleware

import (
	"context"
	"log"

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
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Check authorization
		err := selfAuth(ctx)
		if err != nil {
			log.Printf("Authorization failed: %v\n", err)
			return nil, grpc.Errorf(codes.PermissionDenied, err.Error())
		}
		log.Println("go to handler")
		return handler(ctx, req)
	}
}

func selfAuth(ctx context.Context) error {
	// Extract metadata from context.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("Failed to extract token from metadata: %v\n", ok)
		return grpc.Errorf(codes.Unauthenticated, "metadata cannot be found")
	}

	// Extract token from metadata.
	token := md.Get("token")
	if len(token) == 0 || token[0] != validTokenValue {
		log.Printf("Given token was invalid: %v\n", token)
		return grpc.Errorf(codes.PermissionDenied, "Given token is invalid")
	}

	log.Printf("Authorizaton passed: %v", token)
	return nil
}
