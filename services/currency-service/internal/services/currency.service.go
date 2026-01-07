package services

import (
	"crypto-platform/services/currency-service/internal/models"
	"crypto-platform/services/currency-service/internal/repositories"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type CurrencyService struct {
	repo       repositories.CurrencyRepository
	httpClient *http.Client
	url        string
}

func NewCurrencyService(repo repositories.CurrencyRepository) (*CurrencyService, error) {

	service := &CurrencyService{
		repo: repo,
		httpClient: &http.Client{
			Timeout: 20 * time.Second,
		},
		url: "https://www.tcmb.gov.tr/kurlar/today.xml",
	}

	currencies, err := service.fetchCurrencies()
	if err != nil {
		return nil, fmt.Errorf("unable to fetch currency info from TCMB: %w", err)
	}

	for _, currency := range currencies {
		if currency.Code == "XDR" {
			continue
		}
		if err := service.repo.SetPriceInfo(&currency); err != nil {
			return nil, fmt.Errorf("unable to insert currency info to db: %w", err)
		}
	}

	return service, nil
}

func (s *CurrencyService) GetCurrencies() ([]*models.Currency, error) {
	currencies, err := s.repo.GetPriceInfo()
	if err != nil {
		return nil, err
	}
	return currencies, nil
}

func (s *CurrencyService) RefetchCurrencies() error {
	currencies, err := s.fetchCurrencies()
	if err != nil {
		return fmt.Errorf("unable to fetch currency info from TCMB: %w", err)
	}

	for _, currency := range currencies {
		if currency.Code == "XDR" {
			continue
		}
		if err := s.repo.SetPriceInfo(&currency); err != nil {
			return fmt.Errorf("unable to insert currency info to db: %w", err)
		}
	}

	return nil
}

func (s *CurrencyService) fetchCurrencies() ([]models.Currency, error) {
	res, err := s.httpClient.Get(s.url)
	if err != nil {
		return nil, fmt.Errorf("❌ Hata: Döviz verileri alınamadı: %w", err)
	}
	defer res.Body.Close()

	var tarihDate models.TarihDate
	decoder := xml.NewDecoder(res.Body)

	err = decoder.Decode(&tarihDate)
	if err != nil {
		return nil, fmt.Errorf("❌ Hata: XML verisi decode edilemedi: %w", err)
	}

	return tarihDate.Currencies, nil
}
