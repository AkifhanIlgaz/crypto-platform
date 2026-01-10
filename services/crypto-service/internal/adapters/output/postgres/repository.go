package postgres

import (
	"fmt"

	"github.com/AkifhanIlgaz/services/crypto-service/internal/core/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type cryptoRepository struct {
	db *gorm.DB
}

func NewCryptoRepository(db *gorm.DB) *cryptoRepository {
	return &cryptoRepository{db: db}
}

func (r *cryptoRepository) GetAll() ([]*domain.PriceInfo, error) {
	var entities []*PriceInfoEntity

	if err := r.db.Find(&entities).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch price infos: %w", err)
	}

	return toDomainList(entities), nil
}

func (r *cryptoRepository) Save(priceInfo *domain.PriceInfo) error {
	entity := toEntity(priceInfo)

	err := r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "exchange"}, {Name: "symbol"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"last_updated_at", "price", "high", "low", "open",
			"close", "base_volume", "quote_volume", "change", "change_percent",
		}),
	}).Create(entity).Error

	if err != nil {
		return fmt.Errorf("failed to save price info: %w", err)
	}

	return nil
}

func (r *cryptoRepository) SaveBatch(priceInfos []*domain.PriceInfo) error {
	if len(priceInfos) == 0 {
		return nil
	}

	entities := toEntities(priceInfos)

	err := r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "exchange"}, {Name: "symbol"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"last_updated_at", "price", "high", "low", "open",
			"close", "base_volume", "quote_volume", "change", "change_percent",
		}),
	}).CreateInBatches(entities, 100).Error

	if err != nil {
		return fmt.Errorf("failed to save price infos batch: %w", err)
	}

	return nil
}
