package postgres

import "time"

type PriceInfoEntity struct {
	ID            uint      `gorm:"primaryKey"`
	Exchange      string    `gorm:"column:exchange;uniqueIndex:uniq_exchange_symbol;not null"`
	Symbol        string    `gorm:"column:symbol;uniqueIndex:uniq_exchange_symbol;not null"`
	LastUpdatedAt time.Time `gorm:"column:last_updated_at;not null"`
	Price         float64   `gorm:"column:price;not null"`
	High          float64   `gorm:"column:high"`
	Low           float64   `gorm:"column:low"`
	Open          float64   `gorm:"column:open"`
	Close         float64   `gorm:"column:close"`
	BaseVolume    float64   `gorm:"column:base_volume"`
	QuoteVolume   float64   `gorm:"column:quote_volume"`
	Change        float64   `gorm:"column:change"`
	ChangePercent float64   `gorm:"column:change_percent"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (PriceInfoEntity) TableName() string {
	return "crypto_prices"
}
