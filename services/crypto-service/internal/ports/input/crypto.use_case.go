package input

import "github.com/AkifhanIlgaz/services/crypto-service/internal/core/domain"

type CryptoUseCase interface {
	GetPriceInfos() ([]*domain.PriceInfo, error)
	RefetchPriceInfos() error
}
