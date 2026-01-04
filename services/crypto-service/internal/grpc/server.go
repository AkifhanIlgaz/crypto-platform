package grpc

import (
	"context"
	"crypto-platform/services/crypto-service/internal/services"
	pbCrypto "crypto-platform/shared/proto/crypto"
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
	return nil, nil
}
