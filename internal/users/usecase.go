package users

import (
	"github.com/iamgadfly/go-echo-api/internal/models"
)

type UseCase interface {
	Register(user *models.User) (*models.User, error)
	GetUsers() ([]models.User, error)
	Login(password, email string) (models.UserWithToken, error)
	//Login() (password string, error)
	//Login(ctx context.Context, user *core.User) (core.UserWithToken, error)
	//Update(ctx context.Context, user *core.User) (*core.User, error)
	//FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*core.UsersList, error)
	//GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*core.UsersList, error)
}
