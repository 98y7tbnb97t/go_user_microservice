package grpc

import (
	"net"

	"github.com/98y7tbnb97t/GoMicro/proto/userpb"
	"github.com/98y7tbnb97t/users-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc *user.Service) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcSrv := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))

	return grpcSrv.Serve(lis)
}
