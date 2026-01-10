package postgres

import "time"

type CurrencyEntity struct {
	ID            uint      `gorm:"primaryKey"`
	Code          string    `gorm:"column:code;type:varchar(3);not null;uniqueIndex:idx_currency_code"`
	Price         float64   `gorm:"column:price;not null"`
	LastUpdatedAt time.Time `gorm:"column:last_updated_at;not null"`
}

func (CurrencyEntity) TableName() string {
	return "currencies"
}
