package output

import "github.com/AkifhanIlgaz/services/crypto-service/internal/core/domain"

type CryptoRepository interface {
	GetAll() ([]*domain.PriceInfo, error)
	Save(priceInfo *domain.PriceInfo) error
	SaveBatch(priceInfos []*domain.PriceInfo) error
}
