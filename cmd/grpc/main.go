package main

import (
	"net"

	"github.com/Sahil2k07/gRPC-GO/internal/config"
	"github.com/Sahil2k07/gRPC-GO/internal/database"
	stock "github.com/Sahil2k07/gRPC-GO/internal/generated/stock/proto"
	"github.com/Sahil2k07/gRPC-GO/internal/repository"
	"github.com/Sahil2k07/gRPC-GO/internal/service"
	"github.com/charmbracelet/log"
	"google.golang.org/grpc"
)

func main() {
	configs := config.LoadConfig()
	database.Connect()

	lis, err := net.Listen("tcp", configs.Grpc.GrpcPort)
	if err != nil {
		log.Fatalf("Error starting the TCP Server: %v", err)
	}

	repo := repository.NewInventoryItemRepository()
	stockService := service.NewStockService(repo)

	grpcServer := grpc.NewServer(grpc.Creds(config.GetGrpcCerts()), grpc.UnaryInterceptor(config.AuthInterceptor(configs.Grpc.GrpcToken)))
	stock.RegisterStockServiceServer(grpcServer, stockService)

	log.Infof("gRPC Server starting on port:%v", configs.Grpc.GrpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
