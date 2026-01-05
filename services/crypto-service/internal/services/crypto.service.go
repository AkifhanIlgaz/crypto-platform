package services

import (
	"crypto-platform/services/crypto-service/internal/mappers"
	"crypto-platform/services/crypto-service/internal/models"
	"crypto-platform/services/crypto-service/internal/repositories"
	"crypto-platform/shared/config"
	"fmt"

	ccxt "github.com/ccxt/ccxt/go/v4"
)

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

func (s *CryptoService) GetFromDB(symbol string) ([]*models.PriceInfo, error) {
	priceInfos, err := s.repo.GetPriceInfo(symbol)
	if err != nil {
		return nil, err
	}
	return priceInfos, nil
}

func (s *CryptoService) GetFromExchanges(symbol string) ([]*models.PriceInfo, error) {
	var priceInfos []*models.PriceInfo

	ticker, err := s.kucoin.FetchTicker(symbol)
	if err != nil {
		return nil, fmt.Errorf("[%s] Ticker verisi çekilemedi: %v", "Kucoin", err)
	}

	kucoinPriceInfo := mappers.TickerToPriceInfo(ticker, "Kucoin")

	err = s.repo.SetPriceInfo(kucoinPriceInfo)
	if err != nil {
		return nil, err
	}
	priceInfos = append(priceInfos, kucoinPriceInfo)

	ticker, err = s.binance.FetchTicker(symbol)
	if err != nil {
		return nil, fmt.Errorf("[%s] Ticker verisi çekilemedi: %v", "Binance", err)
	}

	binancePriceInfo := mappers.TickerToPriceInfo(ticker, "Binance")

	err = s.repo.SetPriceInfo(binancePriceInfo)
	if err != nil {
		return nil, err
	}
	priceInfos = append(priceInfos, kucoinPriceInfo)

	return priceInfos, nil
}
