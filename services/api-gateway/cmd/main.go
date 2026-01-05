package main

import (
	"crypto-platform/services/api-gateway/internal/clients"
	"crypto-platform/services/api-gateway/internal/handlers"
	"crypto-platform/shared/config"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	clientManager, err := clients.NewClientManager(&config.CryptoService)
	if err != nil {
		log.Fatal(err)
	}

	defer clientManager.Close()

	logger := logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path} (${latency}) - IP: ${ip} - User-Agent: ${ua}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
		Output:     os.Stdout, //  bir dosyaya veya elasticsearch e yazilabilir
	})
	cryptoHandler := handlers.NewCryptoHandler(clientManager.CryptoClient)

	app := fiber.New(fiber.Config{
		AppName:      "Crypto Platform API Gateway",
		ServerHeader: "Crypto Platform",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	api := app.Group("/api", logger)
	cryptoHandler.RegisterRoutes(api)
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "Route not found",
		})
	})

	addr := fmt.Sprintf(":%v", config.Gateway.Port)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
