package main

import (
	grpcServer "crypto-platform/services/crypto-service/internal/grpc"
	"crypto-platform/services/crypto-service/internal/models"
	"crypto-platform/services/crypto-service/internal/repositories"
	"crypto-platform/services/crypto-service/internal/services"
	"crypto-platform/shared/config"
	"crypto-platform/shared/database"
	pbCrypto "crypto-platform/shared/proto/crypto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, close, err := database.ConnectPostgres(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}()

	db.AutoMigrate(&models.PriceInfo{})

	cryptoRepository := repositories.NewCryptoRepository(db)

	cryptoService, err := services.NewCryptoService(cryptoRepository, cfg.Exchanges)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.CryptoService.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cryptoService.GetFromExchanges("BTC/USDT")

	grpcSrv := grpc.NewServer()
	pbCrypto.RegisterCryptoServiceServer(grpcSrv, grpcServer.NewCryptoServer(cryptoService))

	log.Printf("Driver Service gRPC server listening on :%v", cfg.CryptoService.GRPCPort)
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve:  %v", err)
	}
}
