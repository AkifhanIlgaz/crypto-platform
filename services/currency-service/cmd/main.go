package main

import (
	grpcServer "crypto-platform/services/currency-service/internal/grpc"
	"crypto-platform/services/currency-service/internal/models"
	"crypto-platform/services/currency-service/internal/repositories"
	"crypto-platform/services/currency-service/internal/services"
	"crypto-platform/shared/config"
	"crypto-platform/shared/database"
	pbCurrency "crypto-platform/shared/proto/currency"
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

	err = db.AutoMigrate(&models.Currency{})
	if err != nil {
		log.Fatal(err)
	}

	currencyRepository := repositories.NewCurrencyRepository(db)

	currencyService, err := services.NewCurrencyService(currencyRepository)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.CurrencyService.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pbCurrency.RegisterCurrencyServiceServer(grpcSrv, grpcServer.NewCurrencyServer(currencyService))

	log.Printf("Currency Service gRPC server listening on :%v", cfg.CurrencyService.GRPCPort)
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve:  %v", err)
	}

}
