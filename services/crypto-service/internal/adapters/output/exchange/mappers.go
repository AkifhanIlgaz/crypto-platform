package exchange

import (
	"time"

	"github.com/AkifhanIlgaz/crypto-platform/shared/utils"
	"github.com/AkifhanIlgaz/services/crypto-service/internal/core/domain"
	ccxt "github.com/ccxt/ccxt/go/v4"
)

// CCXT Ticker -> Domain Ticker
func toDomainTicker(ticker ccxt.Ticker, exchange string) *domain.Ticker {
	return &domain.Ticker{
		Exchange:      exchange,
		Symbol:        utils.GetValueOrDefault(ticker.Symbol),
		LastUpdatedAt: time.Now(),
		Last:          utils.GetValueOrDefault(ticker.Last),
		High:          utils.GetValueOrDefault(ticker.High),
		Low:           utils.GetValueOrDefault(ticker.Low),
		Open:          utils.GetValueOrDefault(ticker.Open),
		Close:         utils.GetValueOrDefault(ticker.Close),
		BaseVolume:    utils.GetValueOrDefault(ticker.BaseVolume),
		QuoteVolume:   utils.GetValueOrDefault(ticker.QuoteVolume),
		Change:        utils.GetValueOrDefault(ticker.Change),
		Percentage:    utils.GetValueOrDefault(ticker.Percentage),
	}
}

func toDomainTickers(tickers []ccxt.Ticker, exchange string) []*domain.Ticker {
	result := make([]*domain.Ticker, len(tickers))
	for i, ticker := range tickers {
		result[i] = toDomainTicker(ticker, exchange)
	}
	return result
}
