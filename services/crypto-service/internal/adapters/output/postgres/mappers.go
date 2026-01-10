package postgres

import "github.com/AkifhanIlgaz/services/crypto-service/internal/core/domain"

func (e *PriceInfoEntity) ToDomain() *domain.PriceInfo {
	return &domain.PriceInfo{
		Exchange:      e.Exchange,
		Symbol:        e.Symbol,
		LastUpdatedAt: e.LastUpdatedAt,
		Price:         e.Price,
		High:          e.High,
		Low:           e.Low,
		Open:          e.Open,
		Close:         e.Close,
		BaseVolume:    e.BaseVolume,
		QuoteVolume:   e.QuoteVolume,
		Change:        e.Change,
		ChangePercent: e.ChangePercent,
	}
}

// Domain -> Entity
func toEntity(p *domain.PriceInfo) *PriceInfoEntity {
	return &PriceInfoEntity{
		Exchange:      p.Exchange,
		Symbol:        p.Symbol,
		LastUpdatedAt: p.LastUpdatedAt,
		Price:         p.Price,
		High:          p.High,
		Low:           p.Low,
		Open:          p.Open,
		Close:         p.Close,
		BaseVolume:    p.BaseVolume,
		QuoteVolume:   p.QuoteVolume,
		Change:        p.Change,
		ChangePercent: p.ChangePercent,
	}
}

func toEntities(priceInfos []*domain.PriceInfo) []*PriceInfoEntity {
	entities := make([]*PriceInfoEntity, len(priceInfos))
	for i, p := range priceInfos {
		entities[i] = toEntity(p)
	}
	return entities
}

func toDomainList(entities []*PriceInfoEntity) []*domain.PriceInfo {
	priceInfos := make([]*domain.PriceInfo, len(entities))
	for i, e := range entities {
		priceInfos[i] = e.ToDomain()
	}
	return priceInfos
}
