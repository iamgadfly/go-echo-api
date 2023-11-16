package users

import (
	"github.com/iamgadfly/go-echo-api/internal/models"
)

type UseCase interface {
	Register(ctx map[string]interface{}, user *models.User) (models.UserWithToken, error)
	//Login(ctx context.Context, user *core.User) (core.UserWithToken, error)
	//Update(ctx context.Context, user *core.User) (*core.User, error)
	//FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*core.UsersList, error)
	//GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*core.UsersList, error)
}
