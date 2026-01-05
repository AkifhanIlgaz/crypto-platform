package services

import (
	"crypto-platform/services/crypto-service/internal/mappers"
	"crypto-platform/services/crypto-service/internal/models"
	"crypto-platform/services/crypto-service/internal/repositories"
	"crypto-platform/shared/config"
	"fmt"

	ccxt "github.com/ccxt/ccxt/go/v4"
)

var symbols = []string{"BTC/USDT", "ETH/USDT", "BNB/USDT", "ADA/USDT", "XRP/USDT", "DOT/USDT", "SOL/USDT", "LTC/USDT", "DYDX/USDT", "DYM/USDT"}

type CryptoService struct {
	repo    repositories.CryptoRepository
	kucoin  *ccxt.Kucoin
	binance *ccxt.Binance
}

func NewCryptoService(repo repositories.CryptoRepository, exchanges config.Exchanges) (*CryptoService, error) {
	var kucoinExchange *ccxt.Kucoin
	var binanceExchange *ccxt.Binance

	for name, cfg := range exchanges {
		switch name {
		case "kucoin":
			kucoin := ccxt.NewKucoin(map[string]any{
				"apiKey":          cfg.APIKey,
				"secret":          cfg.APISecret,
				"passphrase":      cfg.Passphrase,
				"enableRateLimit": true,
			})
			if kucoin == nil {
				return nil, fmt.Errorf("[%s] API bağlantısı kurulamadı", "Kucoin")
			}
			kucoinExchange = kucoin

			if _, err := kucoin.LoadMarkets(); err != nil {
				return nil, fmt.Errorf("[%s] Piyasalar yüklenirken hata oluştu: %v", "Kucoin", err)
			}

			tickers, err := kucoin.FetchTickers(ccxt.WithFetchTickersSymbols(symbols))
			if err != nil {
				return nil, fmt.Errorf("[%s] Ticker verisi çekilemedi: %v", "Kucoin", err)
			}

			for _, ticker := range tickers.Tickers {
				repo.SetPriceInfo(mappers.TickerToPriceInfo(ticker, "Kucoin"))
			}
		case "binance":
			binance := ccxt.NewBinance(map[string]any{
				"apiKey":          cfg.APIKey,
				"secret":          cfg.APISecret,
				"enableRateLimit": true,
			})
			if binance == nil {
				return nil, fmt.Errorf("[%s] API bağlantısı kurulamadı", "Binance")
			}
			binanceExchange = binance

			if _, err := binance.LoadMarkets(); err != nil {
				return nil, fmt.Errorf("[%s] Piyasalar yüklenirken hata oluştu: %v", "Binance", err)
			}

			tickers, err := binance.FetchTickers(ccxt.WithFetchTickersSymbols(symbols))
			if err != nil {
				return nil, fmt.Errorf("[%s] Ticker verisi çekilemedi: %v", "Binance", err)
			}

			for _, ticker := range tickers.Tickers {
				repo.SetPriceInfo(mappers.TickerToPriceInfo(ticker, "Binance"))
			}
		default:
			return nil, fmt.Errorf("Geçersiz exchange: %s", name)
		}
	}

	return &CryptoService{
		repo:    repo,
		kucoin:  kucoinExchange,
		binance: binanceExchange,
	}, nil
}

func (s *CryptoService) Get() ([]*models.PriceInfo, error) {
	priceInfos, err := s.repo.GetPriceInfo()
	if err != nil {
		return nil, err
	}
	return priceInfos, nil
}

func (s *CryptoService) Refetch() error {
	tickers, err := s.kucoin.FetchTickers(ccxt.WithFetchTickersSymbols(symbols))
	if err != nil {
		return fmt.Errorf("[%s] Ticker verisi çekilemedi: %v", "Kucoin", err)
	}

	for _, ticker := range tickers.Tickers {
		s.repo.SetPriceInfo(mappers.TickerToPriceInfo(ticker, "Kucoin"))
	}

	tickers, err = s.binance.FetchTickers(ccxt.WithFetchTickersSymbols(symbols))
	if err != nil {
		return fmt.Errorf("[%s] Ticker verisi çekilemedi: %v", "Binance", err)
	}

	for _, ticker := range tickers.Tickers {
		s.repo.SetPriceInfo(mappers.TickerToPriceInfo(ticker, "Binance"))
	}

	return nil
}
