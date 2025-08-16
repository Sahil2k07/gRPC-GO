package config

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type tokenAuth struct {
	token string
}

func (t *tokenAuth) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + t.token,
	}, nil
}

func (t *tokenAuth) RequireTransportSecurity() bool {
	return true
}

func newTokenAuth(token string) (*tokenAuth, error) {
	if len(token) < 32 {
		return nil, fmt.Errorf("API token must be at least 32 characters")
	}
	return &tokenAuth{token: token}, nil
}

var GrpcConn *grpc.ClientConn

func GenerateStockClient(config grpcConfig) {
	creds, err := credentials.NewClientTLSFromFile("certs/server.crt", "")
	if err != nil {
		log.Fatalf("could not load TLS certificate: %v", err)
	}

	auth, err := newTokenAuth(config.GrpcToken)
	if err != nil {
		log.Fatalf("Invalid API token: %v", err)
	}

	conn, err := grpc.NewClient(
		config.GrpcUrl,
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(auth),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	GrpcConn = conn
}

func AuthInterceptor(apiToken string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		tokens := md.Get("Authorization")
		if len(tokens) == 0 || tokens[0] != fmt.Sprintf("Bearer %s", apiToken) {
			return nil, status.Error(codes.Unauthenticated, "invalid API token")
		}

		return handler(ctx, req)
	}
}
