package main

import (
	"fmt"
	"log"
	"net"

	"github.com/AkifhanIlgaz/crypto-platform/shared/config"
	"github.com/AkifhanIlgaz/crypto-platform/shared/database"
	pbCrypto "github.com/AkifhanIlgaz/crypto-platform/shared/proto/crypto"
	grpcServer "github.com/AkifhanIlgaz/services/crypto-service/internal/grpc"
	"github.com/AkifhanIlgaz/services/crypto-service/internal/models"
	"github.com/AkifhanIlgaz/services/crypto-service/internal/repositories"
	"github.com/AkifhanIlgaz/services/crypto-service/internal/services"
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
		fmt.Printf("unable to start crypto service: %s", err.Error())
		return
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.CryptoService.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pbCrypto.RegisterCryptoServiceServer(grpcSrv, grpcServer.NewCryptoServer(cryptoService))

	log.Printf("Crypto Service gRPC server listening on :%v", cfg.CryptoService.GRPCPort)
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve:  %v", err)
	}
}
