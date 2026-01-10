package grpc

import (
	"time"

	pbCrypto "github.com/AkifhanIlgaz/crypto-platform/shared/proto/crypto"
	"github.com/AkifhanIlgaz/services/crypto-service/internal/core/domain"
)

func priceInfoToProto(priceInfo *domain.PriceInfo) *pbCrypto.ExchangePrice {
	return &pbCrypto.ExchangePrice{
		Exchange:      priceInfo.Exchange,
		Price:         priceInfo.Price,
		High:          priceInfo.High,
		Low:           priceInfo.Low,
		Open:          priceInfo.Open,
		Close:         priceInfo.Close,
		BaseVolume:    priceInfo.BaseVolume,
		QuoteVolume:   priceInfo.QuoteVolume,
		Change:        priceInfo.Change,
		ChangePercent: priceInfo.ChangePercent,
		LastUpdatedAt: priceInfo.LastUpdatedAt.Format(time.RFC3339),
	}
}

func priceInfosToExchangePriceListMap(priceInfos []*domain.PriceInfo) map[string]*pbCrypto.ExchangePriceList {
	priceInfoMap := make(map[string]*pbCrypto.ExchangePriceList)

	for _, priceInfo := range priceInfos {
		list, exists := priceInfoMap[priceInfo.Symbol]
		if !exists {
			list = &pbCrypto.ExchangePriceList{
				Exchanges: []*pbCrypto.ExchangePrice{},
			}
			priceInfoMap[priceInfo.Symbol] = list
		}

		list.Exchanges = append(list.Exchanges, priceInfoToProto(priceInfo))
	}

	return priceInfoMap
}
