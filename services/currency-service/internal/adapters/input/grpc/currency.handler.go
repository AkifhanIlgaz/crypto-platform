package grpc

import (
	"context"

	pbCurrency "github.com/AkifhanIlgaz/crypto-platform/shared/proto/currency"
	"github.com/AkifhanIlgaz/services/currency-service/internal/ports/input"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CurrencyHandler struct {
	pbCurrency.UnimplementedCurrencyServiceServer
	useCase input.CurrencyUseCase
}

func NewCurrencyHandler(useCase input.CurrencyUseCase) *CurrencyHandler {
	return &CurrencyHandler{
		useCase: useCase,
	}
}

func (s *CurrencyHandler) GetPriceInfos(ctx context.Context, req *pbCurrency.GetPriceInfosRequest) (*pbCurrency.GetPriceInfosResponse, error) {
	currencies, err := s.useCase.GetCurrencies()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get currencies from database: %v", err)
	}

	return &pbCurrency.GetPriceInfosResponse{
		Currencies: CurrenciesToProto(currencies),
	}, nil
}

func (s *CurrencyHandler) RefetchPriceInfos(ctx context.Context, req *pbCurrency.RefetchPriceInfosRequest) (*pbCurrency.RefetchPriceInfosResponse, error) {
	err := s.useCase.RefetchCurrencies()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get prices from exchanges: %v", err)
	}

	return &pbCurrency.RefetchPriceInfosResponse{
		Success: true,
		Message: "Basarili",
	}, nil
}
