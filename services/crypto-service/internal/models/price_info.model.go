package models

import "time"

type PriceInfo struct {
	ID            uint      `gorm:"primaryKey"`
	Symbol        string    `gorm:"column:symbol;uniqueIndex"`
	LastUpdatedAt time.Time `gorm:"column:last_updated_at"`
	Price         float64   `gorm:"column:price"`
	High          float64   `gorm:"column:high"`
	Low           float64   `gorm:"column:low"`
	Open          float64   `gorm:"column:open"`
	Close         float64   `gorm:"column:close"`
	BaseVolume    float64   `gorm:"column:base_volume"`
	QuoteVolume   float64   `gorm:"column:quote_volume"`
	Change        float64   `gorm:"column:change"`
	ChangePercent float64   `gorm:"column:change_percent"`
}

func (p *PriceInfo) TableName() string {
	return "crypto_prices"
}
