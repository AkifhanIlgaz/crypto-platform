package services

import (
	"fmt"
	"log"

	"github.com/AkifhanIlgaz/services/crypto-service/internal/core/domain"
	"github.com/AkifhanIlgaz/services/crypto-service/internal/ports/output"
)

var symbols = []string{"BTC/USDT", "ETH/USDT", "BNB/USDT", "ADA/USDT", "XRP/USDT", "DOT/USDT", "SOL/USDT", "LTC/USDT", "DYDX/USDT", "SUI/USDT"}

type CryptoService struct {
	repo      output.CryptoRepository
	exchanges []output.ExchangeClient
}

func NewCryptoService(repo output.CryptoRepository, exchanges []output.ExchangeClient) (*CryptoService, error) {
	service := &CryptoService{
		repo:      repo,
		exchanges: exchanges,
	}

	if err := service.RefetchPriceInfos(); err != nil {
		log.Printf("Warning: Initial price fetch failed: %v", err)
	}

	return service, nil
}

func (s *CryptoService) GetPriceInfos() ([]*domain.PriceInfo, error) {
	priceInfos, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return priceInfos, nil
}

func (s *CryptoService) RefetchPriceInfos() error {
	var allPriceInfos []*domain.PriceInfo

	for _, exchange := range s.exchanges {
		tickers, err := exchange.FetchTickers(symbols)
		if err != nil {
			log.Printf("[%s] Failed to fetch tickers: %v", exchange.GetName(), err)
			continue
		}

		for _, ticker := range tickers {
			priceInfo := ticker.ToPriceInfo()
			if err := priceInfo.Validate(); err != nil {
				log.Printf("[%s] Invalid price info for %s: %v",
					exchange.GetName(), ticker.Symbol, err)
				continue
			}
			allPriceInfos = append(allPriceInfos, priceInfo)
		}

		log.Printf("[%s] Successfully fetched %d tickers", exchange.GetName(), len(tickers))
	}

	if len(allPriceInfos) == 0 {
		return fmt.Errorf("no price infos fetched from any exchange")
	}

	if err := s.repo.SaveBatch(allPriceInfos); err != nil {
		return fmt.Errorf("failed to save price infos: %w", err)
	}

	return nil
}

// func (s *CryptoService) Refetch() error {
// 	tickers, err := s.kucoin.FetchTickers(ccxt.WithFetchTickersSymbols(symbols))
// 	if err != nil {
// 		return fmt.Errorf("[%s] Ticker verisi çekilemedi: %v", "Kucoin", err)
// 	}

// 	for _, ticker := range tickers.Tickers {
// 		s.repo.SetPriceInfo(mappers.TickerToPriceInfo(ticker, "Kucoin"))
// 	}

// 	tickers, err = s.binance.FetchTickers(ccxt.WithFetchTickersSymbols(symbols))
// 	if err != nil {
// 		return fmt.Errorf("[%s] Ticker verisi çekilemedi: %v", "Binance", err)
// 	}

// 	for _, ticker := range tickers.Tickers {
// 		s.repo.SetPriceInfo(mappers.TickerToPriceInfo(ticker, "Binance"))
// 	}

// 	tickers, err = s.okx.FetchTickers(ccxt.WithFetchTickersSymbols(symbols))
// 	if err != nil {
// 		return fmt.Errorf("[%s] Ticker verisi çekilemedi: %v", "OKX", err)
// 	}

// 	for _, ticker := range tickers.Tickers {
// 		s.repo.SetPriceInfo(mappers.TickerToPriceInfo(ticker, "OKX"))
// 	}

// 	return nil
// }
