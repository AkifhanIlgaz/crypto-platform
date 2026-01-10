package output

import "github.com/AkifhanIlgaz/services/currency-service/internal/core/domain"

type CurrencyRepository interface {
	GetPriceInfo() ([]*domain.Currency, error)
	SetPriceInfo(priceInfo *domain.Currency) error
}
