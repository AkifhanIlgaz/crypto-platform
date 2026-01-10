package domain

import "time"

type Ticker struct {
	Exchange      string
	Symbol        string
	LastUpdatedAt time.Time
	Last          float64
	High          float64
	Low           float64
	Open          float64
	Close         float64
	BaseVolume    float64
	QuoteVolume   float64
	Change        float64
	Percentage    float64
}

func (t *Ticker) ToPriceInfo() *PriceInfo {
	return &PriceInfo{
		Exchange:      t.Exchange,
		Symbol:        t.Symbol,
		LastUpdatedAt: t.LastUpdatedAt,
		Price:         t.Last,
		High:          t.High,
		Low:           t.Low,
		Open:          t.Open,
		Close:         t.Close,
		BaseVolume:    t.BaseVolume,
		QuoteVolume:   t.QuoteVolume,
		Change:        t.Change,
		ChangePercent: t.Percentage,
	}
}
