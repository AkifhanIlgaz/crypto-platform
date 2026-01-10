package models

import "time"

type Currency struct {
	ID            uint
	LastUpdatedAt time.Time
	Code          string
	Price         float64
}
