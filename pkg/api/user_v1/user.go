package user_v1

import (
	"context"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/iamgadfly/go-echo-api/internal/users/repository"
	"github.com/iamgadfly/go-echo-api/pkg/api/user_v1/proto"
)

// UserServer ..
type UserServer struct {
	proto.UnimplementedUserServer
	cfg  *config.Config
	repo *repository.UsersRepo
}

// NewUserServer ..
func NewUserServer(cfg *config.Config, repo *repository.UsersRepo) *UserServer {
	return &UserServer{
		cfg:  cfg,
		repo: repo,
	}
}

// Get ..
func (u UserServer) Get(ctx context.Context, req *proto.UserIdRequest) (*proto.UserResponse, error) {
	user, err := u.repo.GetById(req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &proto.UserResponse{Data: u.createResponse(user)}, nil
}

// GetUsers ..
func (u UserServer) GetUsers(ctx context.Context, req *proto.EmptyParams) (*proto.UserResponseList, error) {
	list, err := u.repo.GetUsers()
	if err != nil {
		return nil, err
	}
	users := make([]*proto.UserData, len(*list))
	for i, v := range *list {
		users[i] = u.createResponse(&v)
	}
	return &proto.UserResponseList{Data: users}, nil
}

func (u UserServer) createResponse(user *models.User) *proto.UserData {
	return &proto.UserData{
		Id:               uint64(user.ID),
		Name:             user.Name,
		Surname:          user.Surname,
		Email:            user.Email,
		Balance:          uint64(user.Balance),
		IsRegisterSocial: user.IsRegisteredSocial,
		IsPaid:           user.IsPaid,
		Role:             user.Role,
	}
}
