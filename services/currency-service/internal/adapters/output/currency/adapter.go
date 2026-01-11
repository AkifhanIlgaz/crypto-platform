package currency

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/AkifhanIlgaz/services/currency-service/internal/core/domain"
	"github.com/AkifhanIlgaz/services/currency-service/internal/ports/output"
)

type currencyClient struct {
	httpClient *http.Client
	url        string
}

func NewCurrencyClient() output.CurrencyClient {
	return &currencyClient{
		httpClient: http.DefaultClient,
		url:        "https://www.tcmb.gov.tr/kurlar/today.xml",
	}
}

func (c *currencyClient) GetCurrencies() ([]domain.Currency, error) {
	res, err := c.httpClient.Get(c.url)
	if err != nil {
		return nil, fmt.Errorf("❌ Hata: Döviz verileri alınamadı: %w", err)
	}
	defer res.Body.Close()

	var tarihDate domain.TarihDate
	decoder := xml.NewDecoder(res.Body)

	err = decoder.Decode(&tarihDate)
	if err != nil {
		return nil, fmt.Errorf("❌ Hata: XML verisi decode edilemedi: %w", err)
	}

	return tarihDate.Currencies, nil
}
