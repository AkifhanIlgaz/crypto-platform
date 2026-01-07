package services

import (
	"fmt"

	"github.com/AkifhanIlgaz/crypto-platform/shared/config"
	"github.com/AkifhanIlgaz/services/crypto-service/internal/mappers"
	"github.com/AkifhanIlgaz/services/crypto-service/internal/models"
	"github.com/AkifhanIlgaz/services/crypto-service/internal/repositories"
	ccxt "github.com/ccxt/ccxt/go/v4"
)

var symbols = []string{"BTC/USDT", "ETH/USDT", "BNB/USDT", "ADA/USDT", "XRP/USDT", "DOT/USDT", "SOL/USDT", "LTC/USDT", "DYDX/USDT", "SUI/USDT"}

type CryptoService struct {
	repo    repositories.CryptoRepository
	kucoin  *ccxt.Kucoin
	binance *ccxt.Binance
	okx     *ccxt.Okx
}

func NewCryptoService(repo repositories.CryptoRepository, exchanges config.Exchanges) (*CryptoService, error) {
	var kucoinExchange *ccxt.Kucoin
	var binanceExchange *ccxt.Binance
	var okxExchange *ccxt.Okx

	for name, cfg := range exchanges {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("[%s] Panic yakalandı: %v\n", name, r)
				}
			}()

			switch name {
			case "kucoin":
				kucoin := ccxt.NewKucoin(map[string]any{
					"apiKey":          cfg.APIKey,
					"secret":          cfg.APISecret,
					"passphrase":      cfg.Passphrase,
					"enableRateLimit": true,
				})
				if kucoin == nil {
					panic(fmt.Sprintf("[%s] API bağlantısı kurulamadı", "Kucoin"))
				}
				kucoinExchange = kucoin

				if _, err := kucoin.LoadMarkets(); err != nil {

					panic(fmt.Sprintf("[%s] Piyasalar yüklenirken hata oluştu: %v", "Kucoin", err))
				}

				tickers, err := kucoin.FetchTickers(ccxt.WithFetchTickersSymbols(symbols))
				if err != nil {
					panic(fmt.Sprintf("[%s] Ticker verisi çekilemedi: %v", "Kucoin", err))
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
					panic(fmt.Sprintf("[%s] API bağlantısı kurulamadı", "Binance"))
				}
				binanceExchange = binance

				if _, err := binance.LoadMarkets(); err != nil {
					panic(fmt.Sprintf("[%s] Piyasalar yüklenirken hata oluştu: %v", "Binance", err))
				}

				tickers, err := binance.FetchTickers(ccxt.WithFetchTickersSymbols(symbols))
				if err != nil {
					panic(fmt.Sprintf("[%s] Ticker verisi çekilemedi: %v", "Binance", err))
				}

				for _, ticker := range tickers.Tickers {
					repo.SetPriceInfo(mappers.TickerToPriceInfo(ticker, "Binance"))
				}
			case "okx":
				okx := ccxt.NewOkx(map[string]any{
					"apiKey":          cfg.APIKey,
					"secret":          cfg.APISecret,
					"enableRateLimit": true,
				})
				if okx == nil {
					panic(fmt.Sprintf("[%s] API bağlantısı kurulamadı", "OKX"))
				}
				okxExchange = okx

				if _, err := okx.LoadMarkets(); err != nil {
					panic(fmt.Sprintf("[%s] Piyasalar yüklenirken hata oluştu: %v", "OKX", err))
				}

				tickers, err := okx.FetchTickers(ccxt.WithFetchTickersSymbols(symbols))
				if err != nil {
					panic(fmt.Sprintf("[%s] Ticker verisi çekilemedi: %v", "OKX", err))
				}

				for _, ticker := range tickers.Tickers {
					repo.SetPriceInfo(mappers.TickerToPriceInfo(ticker, "OKX"))
				}
			default:
				panic(fmt.Sprintf("Geçersiz exchange: %s", name))
			}
		}()
	}

	return &CryptoService{
		repo:    repo,
		kucoin:  kucoinExchange,
		binance: binanceExchange,
		okx:     okxExchange,
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

	tickers, err = s.okx.FetchTickers(ccxt.WithFetchTickersSymbols(symbols))
	if err != nil {
		return fmt.Errorf("[%s] Ticker verisi çekilemedi: %v", "OKX", err)
	}

	for _, ticker := range tickers.Tickers {
		s.repo.SetPriceInfo(mappers.TickerToPriceInfo(ticker, "OKX"))
	}

	return nil
}
