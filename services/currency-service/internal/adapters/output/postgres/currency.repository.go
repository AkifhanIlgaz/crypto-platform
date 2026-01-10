package postgres

import (
	"github.com/AkifhanIlgaz/services/currency-service/internal/core/domain"
	"github.com/AkifhanIlgaz/services/currency-service/internal/ports/output"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type currencyRepository struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) output.CurrencyRepository {
	return currencyRepository{db: db}
}

func (r currencyRepository) GetPriceInfo() ([]*domain.Currency, error) {
	var currencies []*domain.Currency
	if err := r.db.Find(&currencies).Error; err != nil {
		return nil, err
	}

	return currencies, nil
}

func (r currencyRepository) SetPriceInfo(priceInfo *domain.Currency) error {
	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "code"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"last_updated_at",
			"price",
		}),
	}).Create(priceInfo).Error
}
