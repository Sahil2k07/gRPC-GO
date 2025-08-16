package config

import "os"

type grpcConfig struct {
	GrpcPort  string `toml:"grpc_port"`
	GrpcUrl   string `toml:"grpc_url"`
	GrpcToken string `toml:"grpc_token"`
}

func loadProdGrpcConfig() grpcConfig {
	return grpcConfig{
		GrpcPort:  os.Getenv("GRPC_PORT"),
		GrpcUrl:   os.Getenv("GRPC_URL"),
		GrpcToken: os.Getenv("GRPC_TOKEN"),
	}
}

func GetGrpcConfig() grpcConfig {
	return grpcConfig{
		GrpcPort:  globalConfig.Grpc.GrpcPort,
		GrpcUrl:   globalConfig.Grpc.GrpcUrl,
		GrpcToken: globalConfig.Grpc.GrpcToken,
	}
}
