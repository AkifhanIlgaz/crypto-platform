package main

import (
	"crypto-platform/services/crypto-service/internal/models"
	"crypto-platform/shared/config"
	"crypto-platform/shared/database"
	"log"
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

	// Create repositories and services
	// listen grpc port
	// yeni grpc server olustur
	// register cryptoServiceServer
	// grpc serve

}
