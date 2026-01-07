package repositories

import (
	"github.com/AkifhanIlgaz/services/currency-service/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CurrencyRepository interface {
	GetPriceInfo() ([]*models.Currency, error)
	SetPriceInfo(priceInfo *models.Currency) error
}

type currencyRepository struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *currencyRepository {
	return &currencyRepository{db: db}
}

func (r *currencyRepository) GetPriceInfo() ([]*models.Currency, error) {
	var currencies []*models.Currency
	if err := r.db.Find(&currencies).Error; err != nil {
		return nil, err
	}

	return currencies, nil
}

func (r *currencyRepository) SetPriceInfo(priceInfo *models.Currency) error {
	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "code"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"last_updated_at",
			"price",
		}),
	}).Create(priceInfo).Error
}
