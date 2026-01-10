package grpc

import (
	pbCurrency "github.com/AkifhanIlgaz/crypto-platform/shared/proto/currency"
	"github.com/AkifhanIlgaz/services/currency-service/internal/core/domain"
)

func currencyToProto(currency *domain.Currency) *pbCurrency.Currency {
	return &pbCurrency.Currency{
		Code:  currency.Code,
		Price: currency.Price,
	}
}

func CurrenciesToProto(currencies []*domain.Currency) []*pbCurrency.Currency {
	var result []*pbCurrency.Currency
	for _, currency := range currencies {
		result = append(result, currencyToProto(currency))
	}
	return result
}
