package users

import (
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/labstack/echo/v4"
)

type UseCase interface {
	Register(c echo.Context, user *models.User) (*models.User, error)
	GetUsers() ([]models.User, error)
	Login(u models.UserLogin) (models.UserWithToken, error)
	GetByID(id interface{}) (*models.User, error)
	//Login() (password string, error)
	//Login(ctx context.Context, user *core.User) (core.UserWithToken, error)
	//Update(ctx context.Context, user *core.User) (*core.User, error)
	//FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*core.UsersList, error)
	//GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*core.UsersList, error)
}
