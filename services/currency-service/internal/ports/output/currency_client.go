package output

import (
	"github.com/AkifhanIlgaz/services/currency-service/internal/core/domain"
)

type CurrencyClient interface {
	GetCurrencies() ([]domain.Currency, error)
}
