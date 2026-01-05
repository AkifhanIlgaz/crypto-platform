package repositories

import (
	"crypto-platform/services/crypto-service/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CryptoRepository interface {
	GetPriceInfo(symbol string) ([]*models.PriceInfo, error)
	SetPriceInfo(priceInfo *models.PriceInfo) error
}

type cryptoRepository struct {
	db *gorm.DB
}

func NewCryptoRepository(db *gorm.DB) *cryptoRepository {
	return &cryptoRepository{db: db}
}

func (r *cryptoRepository) GetPriceInfo(symbol string) ([]*models.PriceInfo, error) {
	var priceInfo []*models.PriceInfo
	if err := r.db.Find(&priceInfo, symbol).Error; err != nil {
		return nil, err
	}
	return priceInfo, nil
}

func (r *cryptoRepository) SetPriceInfo(priceInfo *models.PriceInfo) error {
	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "exchange"}, {Name: "symbol"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"last_updated_at",
			"price",
			"high",
			"low",
			"open",
			"close",
			"base_volume",
			"quote_volume",
			"change",
			"change_percent",
		}),
	}).Create(priceInfo).Error
}
