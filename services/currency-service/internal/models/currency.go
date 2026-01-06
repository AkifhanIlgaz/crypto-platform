package models

import (
	"encoding/xml"
	"time"
)

type TarihDate struct {
	XMLName    xml.Name   `xml:"Tarih_Date"`
	Date       string     `xml:"Date,attr"`
	Currencies []Currency `xml:"Currency"`
}

type Currency struct {
	ID            uint      `gorm:"primaryKey"`
	LastUpdatedAt time.Time `gorm:"column:last_updated_at;type:datetime;not null"`
	Code          string    `xml:"CurrencyCode,attr" gorm:"column:code;type:varchar(3);not null"`
	Price  float64   `xml:"ForexSelling" gorm:"column:price;not null"`
}
