package main

import (
	"fmt"
	"log"
	"net"

	"github.com/AkifhanIlgaz/crypto-platform/shared/config"
	"github.com/AkifhanIlgaz/crypto-platform/shared/database"
	pbCurrency "github.com/AkifhanIlgaz/crypto-platform/shared/proto/currency"
	grpcHandler "github.com/AkifhanIlgaz/services/currency-service/internal/adapters/input/grpc"
	"github.com/AkifhanIlgaz/services/currency-service/internal/adapters/output/postgres"
	"github.com/AkifhanIlgaz/services/currency-service/internal/core/services"

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

	err = db.AutoMigrate(&postgres.CurrencyEntity{})
	if err != nil {
		log.Fatal(err)
	}

	currencyRepository := postgres.NewCurrencyRepository(db)

	currencyService, err := services.NewCurrencyService(currencyRepository)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.CurrencyService.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pbCurrency.RegisterCurrencyServiceServer(grpcSrv, grpcHandler.NewCurrencyHandler(currencyService))

	log.Printf("Currency Service gRPC server listening on :%v", cfg.CurrencyService.GRPCPort)
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve:  %v", err)
	}

}
