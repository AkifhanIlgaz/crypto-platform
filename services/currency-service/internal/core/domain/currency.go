package domain

import (
	"encoding/xml"
	"fmt"
	"time"
)

type TarihDate struct {
	XMLName    xml.Name   `xml:"Tarih_Date"`
	Date       string     `xml:"Date,attr"`
	Currencies []Currency `xml:"Currency"`
}
type Currency struct {
	Code          string
	Price         float64
	LastUpdatedAt time.Time
}

func (c *Currency) Validate() error {
	if c.Code == "" {
		return fmt.Errorf("currency code is required")
	}

	if len(c.Code) != 3 {
		return fmt.Errorf("currency code must be 3 characters")
	}

	if c.Price < 0 {
		return fmt.Errorf("price cannot be negative")
	}

	return nil
}
