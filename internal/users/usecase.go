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
}
