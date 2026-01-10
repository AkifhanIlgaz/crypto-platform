package main

import (
	"fmt"
	"log"
	"net"

	"github.com/AkifhanIlgaz/crypto-platform/shared/config"
	"github.com/AkifhanIlgaz/crypto-platform/shared/database"
	pbCrypto "github.com/AkifhanIlgaz/crypto-platform/shared/proto/crypto"
	grpcHandler "github.com/AkifhanIlgaz/services/crypto-service/internal/adapters/input/grpc"
	exchangeAdapter "github.com/AkifhanIlgaz/services/crypto-service/internal/adapters/output/exchange"
	"github.com/AkifhanIlgaz/services/crypto-service/internal/adapters/output/postgres"
	"github.com/AkifhanIlgaz/services/crypto-service/internal/core/services"
	"github.com/AkifhanIlgaz/services/crypto-service/internal/ports/output"

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

	db.AutoMigrate(&postgres.PriceInfoEntity{})

	cryptoRepository := postgres.NewCryptoRepository(db)

	exchanges := initializeExchanges(cfg.Exchanges)

	cryptoService, err := services.NewCryptoService(cryptoRepository, exchanges)
	if err != nil {
		fmt.Printf("unable to start crypto service: %s", err.Error())
		return
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.CryptoService.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pbCrypto.RegisterCryptoServiceServer(grpcSrv, grpcHandler.NewCryptoHandler(cryptoService))

	log.Printf("Crypto Service gRPC server listening on :%v", cfg.CryptoService.GRPCPort)
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve:  %v", err)
	}
}

func initializeExchanges(exchangeConfigs config.Exchanges) []output.ExchangeClient {
	var exchanges []output.ExchangeClient

	for name, cfg := range exchangeConfigs {
		log.Printf("Initializing exchange: %s", name)

		adapter, err := exchangeAdapter.NewCCXTAdapter(name, map[string]any{
			"apiKey":          cfg.APIKey,
			"secret":          cfg.APISecret,
			"passphrase":      cfg.Passphrase,
			"enableRateLimit": true,
		})

		if err != nil {
			log.Printf("⚠️  Failed to initialize %s: %v", name, err)
			continue
		}

		exchanges = append(exchanges, adapter)
		log.Printf("✅ Successfully initialized %s", name)
	}

	return exchanges
}
