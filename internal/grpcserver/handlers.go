package grpcserver

import (
	"github.com/iamgadfly/go-echo-api/internal/users/repository"
	"github.com/iamgadfly/go-echo-api/pkg/api/user_v1"
	"github.com/iamgadfly/go-echo-api/pkg/api/user_v1/proto"
	"google.golang.org/grpc"
)

// MapServices ..
func (s *ServerGRPC) MapServices(server *grpc.Server) {
	userRepo := repository.NewUsersRepository(s.db)
	service := user_v1.NewUserServer(s.cfg, userRepo)
	proto.RegisterUserServer(server, service)
}
