package grpc

import (
	"context"
	"log"

	pbCrypto "github.com/AkifhanIlgaz/crypto-platform/shared/proto/crypto"
	"github.com/AkifhanIlgaz/services/crypto-service/internal/ports/input"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CryptoHandler struct {
	pbCrypto.UnimplementedCryptoServiceServer
	useCase input.CryptoUseCase
}

func NewCryptoHandler(useCase input.CryptoUseCase) *CryptoHandler {
	return &CryptoHandler{
		useCase: useCase,
	}
}

func (h *CryptoHandler) GetPriceInfos(
	ctx context.Context,
	req *pbCrypto.GetPriceInfosRequest,
) (*pbCrypto.GetPriceInfosResponse, error) {
	log.Printf("GetPriceInfos request received")

	priceInfos, err := h.useCase.GetPriceInfos()
	if err != nil {
		log.Printf("Failed to get price infos: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get prices: %v", err)
	}

	return &pbCrypto.GetPriceInfosResponse{
		Prices: priceInfosToExchangePriceListMap(priceInfos),
	}, nil
}

func (h *CryptoHandler) RefetchPriceInfos(
	ctx context.Context,
	req *pbCrypto.RefetchPriceInfosRequest,
) (*pbCrypto.RefetchPriceInfosResponse, error) {
	log.Printf("RefetchPriceInfos request received")

	err := h.useCase.RefetchPriceInfos()
	if err != nil {
		log.Printf("Failed to refetch price infos: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to refetch prices: %v", err)
	}

	return &pbCrypto.RefetchPriceInfosResponse{
		Success: true,
		Message: "Price infos successfully refetched",
	}, nil
}

// package grpc

// import (
// 	"context"

// 	pbCrypto "github.com/AkifhanIlgaz/crypto-platform/shared/proto/crypto"
// 	"github.com/AkifhanIlgaz/services/crypto-service/internal/core/mappers"
// 	"github.com/AkifhanIlgaz/services/crypto-service/internal/core/services"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// type CryptoHandler struct {
// 	pbCrypto.UnimplementedCryptoServiceServer
// 	service *services.CryptoService
// }

// func NewCryptoHandler(service *services.CryptoService) *CryptoHandler {
// 	return &CryptoHandler{
// 		service: service,
// 	}
// }

// func (s *CryptoHandler) GetPriceInfos(ctx context.Context, req *pbCrypto.GetPriceInfosRequest) (*pbCrypto.GetPriceInfosResponse, error) {
// 	prices, err := s.service.Get()
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to get prices from database: %v", err)
// 	}

// 	return &pbCrypto.GetPriceInfosResponse{
// 		Prices: mappers.PriceInfosToExchangePriceListMap(prices),
// 	}, nil
// }

// func (s *CryptoHandler) RefetchPriceInfos(ctx context.Context, req *pbCrypto.RefetchPriceInfosRequest) (*pbCrypto.RefetchPriceInfosResponse, error) {
// 	err := s.service.Refetch()
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to get prices from exchanges: %v", err)
// 	}

// 	return &pbCrypto.RefetchPriceInfosResponse{
// 		Success: true,
// 		Message: "Basarili",
// 	}, nil
// }
