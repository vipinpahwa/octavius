package registration

import (
	"context"
	"octavius/internal/pkg/protofiles"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client interface {
	ConnectClient(cpHost string, connectionTimeOut time.Duration) error
	Register(request *protofiles.RegisterRequest) (*protofiles.RegisterResponse, error)
}

type GrpcClient struct {
	client                protofiles.RegistrationServiceClient
	connectionTimeoutSecs time.Duration
}

func (g *GrpcClient) ConnectClient(cpHost string, connectionTimeOut time.Duration) error {
	conn, err := grpc.Dial(cpHost, grpc.WithInsecure())
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	grpcClient := protofiles.NewRegistrationServiceClient(conn)
	g.client = grpcClient
	g.connectionTimeoutSecs = connectionTimeOut
	return nil
}

func (g *GrpcClient) Register(request *protofiles.RegisterRequest) (*protofiles.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()

	return g.client.Register(ctx, request)
}
