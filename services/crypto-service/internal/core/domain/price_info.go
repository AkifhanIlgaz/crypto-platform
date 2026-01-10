package domain

import (
	"fmt"
	"time"
)

type PriceInfo struct {
	Exchange      string
	Symbol        string
	LastUpdatedAt time.Time
	Price         float64
	High          float64
	Low           float64
	Open          float64
	Close         float64
	BaseVolume    float64
	QuoteVolume   float64
	Change        float64
	ChangePercent float64
}

func (p *PriceInfo) Validate() error {
	if p.Exchange == "" || p.Symbol == "" {
		return fmt.Errorf("exchange and symbol are required")
	}
	if p.Price < 0 {
		return fmt.Errorf("price must be positive")
	}
	return nil
}
