package handlers

import (
	"context"
	"time"

	"github.com/AkifhanIlgaz/crypto-platform/shared/response"

	pbCurrency "github.com/AkifhanIlgaz/crypto-platform/shared/proto/currency"
	"github.com/gofiber/fiber/v2"
)

type CurrencyHandler struct {
	client pbCurrency.CurrencyServiceClient
}

func NewCurrencyHandler(client pbCurrency.CurrencyServiceClient) *CurrencyHandler {
	return &CurrencyHandler{client: client}
}

func (h *CurrencyHandler) RegisterRoutes(router fiber.Router) {
	cryptoRoute := router.Group("/currency")
	cryptoRoute.Get("/prices", h.GetCryptoPrices)
	cryptoRoute.Post("/refetch", h.Refetch)
}

func (h *CurrencyHandler) GetCryptoPrices(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := h.client.GetPriceInfos(ctx, &pbCurrency.GetPriceInfosRequest{})
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, res, "Coin verileri başarı ile getirildi !")
}

func (h *CurrencyHandler) Refetch(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := h.client.RefetchPriceInfos(ctx, &pbCurrency.RefetchPriceInfosRequest{})
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, res, "Coin verileri başarı ile yenilendi !")
}
