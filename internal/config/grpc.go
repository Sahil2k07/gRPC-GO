package config

import (
	"os"

	"github.com/charmbracelet/log"
	"google.golang.org/grpc/credentials"
)

type grpcConfig struct {
	GrpcPort  string `toml:"grpc_port"`
	GrpcUrl   string `toml:"grpc_url"`
	GrpcToken string `toml:"grpc_token"`
	GrpcCerts credentials.TransportCredentials
}

func loadGrpcCerts() credentials.TransportCredentials {
	creds, err := credentials.NewClientTLSFromFile("certs/server.crt", "")
	if err != nil {
		log.Fatalf("could not load TLS certificate: %v", err)
	}

	return creds
}

func GetGrpcCerts() credentials.TransportCredentials {
	return globalConfig.Grpc.GrpcCerts
}

func loadProdGrpcConfig() grpcConfig {
	return grpcConfig{
		GrpcPort:  os.Getenv("GRPC_PORT"),
		GrpcUrl:   os.Getenv("GRPC_URL"),
		GrpcToken: os.Getenv("GRPC_TOKEN"),
		GrpcCerts: loadGrpcCerts(),
	}
}

func GetGrpcConfig() grpcConfig {
	return grpcConfig{
		GrpcPort:  globalConfig.Grpc.GrpcPort,
		GrpcUrl:   globalConfig.Grpc.GrpcUrl,
		GrpcToken: globalConfig.Grpc.GrpcToken,
		GrpcCerts: globalConfig.Grpc.GrpcCerts,
	}
}
