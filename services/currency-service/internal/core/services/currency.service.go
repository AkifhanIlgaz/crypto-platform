package services

import (
	"fmt"

	"github.com/AkifhanIlgaz/services/currency-service/internal/core/domain"
	"github.com/AkifhanIlgaz/services/currency-service/internal/ports/output"
)

type CurrencyService struct {
	repo           output.CurrencyRepository
	currencyClient output.CurrencyClient
}

func NewCurrencyService(repo output.CurrencyRepository, currencyClient output.CurrencyClient) (*CurrencyService, error) {

	service := &CurrencyService{
		repo:           repo,
		currencyClient: currencyClient,
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

func (s *CurrencyService) GetCurrencies() ([]*domain.Currency, error) {
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

func (s *CurrencyService) fetchCurrencies() ([]domain.Currency, error) {
	currencies, err := s.currencyClient.GetCurrencies()
	if err != nil {
		return nil, fmt.Errorf("❌ Hata: Döviz verileri alınamadı: %w", err)
	}

	return currencies, nil
}
