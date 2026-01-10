package exchange

import (
	"fmt"

	"github.com/AkifhanIlgaz/services/crypto-service/internal/core/domain"
	ccxt "github.com/ccxt/ccxt/go/v4"
)

type ccxtExchangeInterface interface {
	FetchTickers(options ...ccxt.FetchTickersOptions) (ccxt.Tickers, error)
	LoadMarkets(params ...interface{}) (map[string]ccxt.MarketInterface, error)
}

type ccxtAdapter struct {
	exchange ccxtExchangeInterface
	name     string
}

func NewCCXTAdapter(name string, config map[string]any) (*ccxtAdapter, error) {
	var exchange ccxtExchangeInterface

	switch name {
	case "kucoin":
		exchange = ccxt.NewKucoin(config)
	case "binance":
		exchange = ccxt.NewBinance(config)
	case "okx":
		exchange = ccxt.NewOkx(config)
	default:
		return nil, fmt.Errorf("unsupported exchange: %s", name)
	}

	if exchange == nil {
		return nil, fmt.Errorf("failed to create exchange client for %s", name)
	}

	adapter := &ccxtAdapter{
		exchange: exchange,
		name:     name,
	}

	if err := adapter.LoadMarkets(); err != nil {
		return nil, fmt.Errorf("failed to load markets for %s: %w", name, err)
	}

	return adapter, nil
}

func (a *ccxtAdapter) FetchTickers(symbols []string) ([]*domain.Ticker, error) {
	tickers, err := a.exchange.FetchTickers(ccxt.WithFetchTickersSymbols(symbols))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tickers: %w", err)
	}

	tickerList := make([]ccxt.Ticker, len(tickers.Tickers))
	for _, ticker := range tickers.Tickers {
		tickerList = append(tickerList, ticker)
	}

	return toDomainTickers(tickerList, a.name), nil
}

func (a *ccxtAdapter) LoadMarkets() error {
	_, err := a.exchange.LoadMarkets()
	if err != nil {
		return fmt.Errorf("failed to load markets: %w", err)
	}
	return nil
}

func (a *ccxtAdapter) GetName() string {
	return a.name
}
