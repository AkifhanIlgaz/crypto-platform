package grpc

import (
	"context"

	pbCurrency "github.com/AkifhanIlgaz/crypto-platform/shared/proto/currency"
	"github.com/AkifhanIlgaz/services/currency-service/internal/mappers"
	"github.com/AkifhanIlgaz/services/currency-service/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CurrencyServer struct {
	pbCurrency.UnimplementedCurrencyServiceServer
	service *services.CurrencyService
}

func NewCurrencyServer(service *services.CurrencyService) *CurrencyServer {
	return &CurrencyServer{
		service: service,
	}
}

func (s *CurrencyServer) GetPriceInfos(ctx context.Context, req *pbCurrency.GetPriceInfosRequest) (*pbCurrency.GetPriceInfosResponse, error) {
	currencies, err := s.service.GetCurrencies()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get currencies from database: %v", err)
	}

	return &pbCurrency.GetPriceInfosResponse{
		Currencies: mappers.CurrenciesToProto(currencies),
	}, nil
}

func (s *CurrencyServer) RefetchPriceInfos(ctx context.Context, req *pbCurrency.RefetchPriceInfosRequest) (*pbCurrency.RefetchPriceInfosResponse, error) {
	err := s.service.RefetchCurrencies()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get prices from exchanges: %v", err)
	}

	return &pbCurrency.RefetchPriceInfosResponse{
		Success: true,
		Message: "Basarili",
	}, nil
}
