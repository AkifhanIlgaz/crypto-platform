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

func (s *CryptoServer) GetPriceInfo(ctx context.Context, req *pbCrypto.GetPriceInfoRequest) (*pbCrypto.GetPriceInfoResponse, error) {
	if req.IsRefetch {
		prices, err := s.service.GetFromExchanges(req.Symbol)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get prices from exchanges: %v", err)
		}

		return &pbCrypto.GetPriceInfoResponse{
			PriceInfos: mappers.PriceInfosToProto(prices),
		}, nil
	}

	prices, err := s.service.GetFromDB(req.Symbol)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get prices from database: %v", err)
	}

	return &pbCrypto.GetPriceInfoResponse{
		PriceInfos: mappers.PriceInfosToProto(prices),
	}, nil
}
