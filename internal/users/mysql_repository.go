package users

import "github.com/iamgadfly/go-echo-api/internal/models"

type Repository interface {
	Create(user *models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	GetByID(id interface{}) (*models.User, error)
	Login(u models.UserLogin) (models.User, error)
}
