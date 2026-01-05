package clients

import (
	"crypto-platform/shared/config"
	pbCrypto "crypto-platform/shared/proto/crypto"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type ClientManager struct {
	CryptoClient pbCrypto.CryptoServiceClient

	cryptoConn *grpc.ClientConn
}

func NewClientManager(cryptoConfig *config.Service) (*ClientManager, error) {
	addr := fmt.Sprintf("%s:%s", cryptoConfig.Name, cryptoConfig.GRPCPort)
	cryptoConn, err := createGRPCConnection(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to crypto service: %w", err)
	}

	return &ClientManager{
		CryptoClient: pbCrypto.NewCryptoServiceClient(cryptoConn),
		cryptoConn:   cryptoConn,
	}, nil
}

func createGRPCConnection(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(10*1024*1024),
			grpc.MaxCallSendMsgSize(10*1024*1024),
		),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             3 * time.Second,
			PermitWithoutStream: true,
		}))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (cm *ClientManager) Close() {
	if cm.cryptoConn != nil {
		cm.cryptoConn.Close()
	}
}
