package mappers

import (
	"crypto-platform/services/crypto-service/internal/models"
	pbCrypto "crypto-platform/shared/proto/crypto"
	"crypto-platform/shared/utils"
	"time"

	ccxt "github.com/ccxt/ccxt/go/v4"
)

func TickerToPriceInfo(ticker ccxt.Ticker, exchange string) *models.PriceInfo {
	return &models.PriceInfo{
		Symbol:        utils.GetValueOrDefault(ticker.Symbol),
		High:          utils.GetValueOrDefault(ticker.High),
		Low:           utils.GetValueOrDefault(ticker.Low),
		Open:          utils.GetValueOrDefault(ticker.Open),
		Close:         utils.GetValueOrDefault(ticker.Close),
		BaseVolume:    utils.GetValueOrDefault(ticker.BaseVolume),
		QuoteVolume:   utils.GetValueOrDefault(ticker.QuoteVolume),
		Price:         utils.GetValueOrDefault(ticker.Last),
		Change:        utils.GetValueOrDefault(ticker.Change),
		ChangePercent: utils.GetValueOrDefault(ticker.Percentage),
		LastUpdatedAt: time.Now(),
		Exchange:      exchange,
	}
}

func priceInfoToProto(priceInfo *models.PriceInfo) *pbCrypto.ExchangePrice {
	return &pbCrypto.ExchangePrice{
		High:          priceInfo.High,
		Low:           priceInfo.Low,
		Open:          priceInfo.Open,
		Close:         priceInfo.Close,
		BaseVolume:    priceInfo.BaseVolume,
		QuoteVolume:   priceInfo.QuoteVolume,
		Price:         priceInfo.Price,
		Change:        priceInfo.Change,
		ChangePercent: priceInfo.ChangePercent,
		LastUpdatedAt: priceInfo.LastUpdatedAt.Format(time.RFC3339),
		Exchange:      priceInfo.Exchange,
	}
}

func PriceInfosToExchangePriceListMap(priceInfos []*models.PriceInfo) map[string]*pbCrypto.ExchangePriceList {
	priceInfoMap := make(map[string]*pbCrypto.ExchangePriceList)
	for _, priceInfo := range priceInfos {
		list, ok := priceInfoMap[priceInfo.Symbol]
		if !ok {
			list = &pbCrypto.ExchangePriceList{
				Exchanges: []*pbCrypto.ExchangePrice{priceInfoToProto(priceInfo)},
			}
			priceInfoMap[priceInfo.Symbol] = list
			continue
		}

		list.Exchanges = append(list.Exchanges, priceInfoToProto(priceInfo))
		priceInfoMap[priceInfo.Symbol] = list
	}
	return priceInfoMap
}
