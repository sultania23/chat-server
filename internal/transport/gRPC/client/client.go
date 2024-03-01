package client

import (
	"github.com/tuxoo/idler/internal/transport/gRPC/api"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	grpcConnection *grpc.ClientConn
	MailSender     api.MailSenderServiceClient
}

func NewGrpcClient(grpcConnection *grpc.ClientConn) *GrpcClient {
	return &GrpcClient{
		grpcConnection: grpcConnection,
		MailSender:     api.NewMailSenderServiceClient(grpcConnection),
	}
}
