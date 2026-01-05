package handlers

import (
	"context"
	pbCrypto "crypto-platform/shared/proto/crypto"
	"crypto-platform/shared/response"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CryptoHandler struct {
	client pbCrypto.CryptoServiceClient
}

func NewCryptoHandler(client pbCrypto.CryptoServiceClient) *CryptoHandler {
	return &CryptoHandler{client: client}
}

func (h *CryptoHandler) RegisterRoutes(router fiber.Router) {
	cryptoRoute := router.Group("/crypto")
	cryptoRoute.Get("/prices", h.GetCryptoPrices)
	cryptoRoute.Post("/refetch", h.Refetch)
}

func (h *CryptoHandler) GetCryptoPrices(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := h.client.GetPriceInfos(ctx, &pbCrypto.GetPriceInfosRequest{})
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, res, "Coin verileri başarı ile getirildi !")
}

func (h *CryptoHandler) Refetch(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := h.client.RefetchPriceInfos(ctx, &pbCrypto.RefetchPriceInfosRequest{})
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, res, "Coin verileri başarı ile yenilendi !")
}
