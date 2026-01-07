package mappers

import (
	pbCurrency "github.com/AkifhanIlgaz/crypto-platform/shared/proto/currency"
	"github.com/AkifhanIlgaz/services/currency-service/internal/models"
)

func currencyToProto(currency *models.Currency) *pbCurrency.Currency {
	return &pbCurrency.Currency{
		Code:  currency.Code,
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
