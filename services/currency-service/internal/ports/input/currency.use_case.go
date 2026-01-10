package input

import (
	"github.com/AkifhanIlgaz/services/currency-service/internal/core/domain"
)

type CurrencyUseCase interface {
	GetCurrencies() ([]*domain.Currency, error)
	RefetchCurrencies() error
}
