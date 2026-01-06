package mappers

import (
	"crypto-platform/services/currency-service/internal/models"
	pbCurrency "crypto-platform/shared/proto/currency"
	"time"
)

func currencyToProto(currency *models.Currency) *pbCurrency.Currency {
	return &pbCurrency.Currency{
		Code:  currency.Code,
		Date:  currency.LastUpdatedAt.Format(time.RFC3339),
		Price: currency.Price,
	}
}

func CurrenciesToProto(currencies []*models.Currency) []*pbCurrency.Currency {
	var result []*pbCurrency.Currency
	for _, currency := range currencies {
		result = append(result, currencyToProto(currency))
	}
	return result
}
