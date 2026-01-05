package grpc

import (
	"context"
	"crypto-platform/services/crypto-service/internal/mappers"
	"crypto-platform/services/crypto-service/internal/services"
	pbCrypto "crypto-platform/shared/proto/crypto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CryptoServer struct {
	pbCrypto.UnimplementedCryptoServiceServer
	service *services.CryptoService
}

func NewCryptoServer(service *services.CryptoService) *CryptoServer {
	return &CryptoServer{
		service: service,
	}
}

func (s *CryptoServer) GetPriceInfos(ctx context.Context, req *pbCrypto.GetPriceInfosRequest) (*pbCrypto.GetPriceInfosResponse, error) {

	prices, err := s.service.Get()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get prices from database: %v", err)
	}

	return &pbCrypto.GetPriceInfosResponse{
		Prices: mappers.PriceInfosToExchangePriceListMap(prices),
	}, nil
}

func (s *CryptoServer) RefetchPriceInfos(ctx context.Context, req *pbCrypto.RefetchPriceInfosRequest) (*pbCrypto.RefetchPriceInfosResponse, error) {
	err := s.service.Refetch()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get prices from exchanges: %v", err)
	}

	return &pbCrypto.RefetchPriceInfosResponse{
		Success: true,
		Message: "Basarili",
	}, nil
}
