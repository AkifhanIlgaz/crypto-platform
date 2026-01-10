package output

import "github.com/AkifhanIlgaz/services/crypto-service/internal/core/domain"

type ExchangeClient interface {
	FetchTickers(symbols []string) ([]*domain.Ticker, error)
	LoadMarkets() error
	GetName() string
}
