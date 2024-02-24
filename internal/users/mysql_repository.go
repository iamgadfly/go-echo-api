package users

import "github.com/iamgadfly/go-echo-api/internal/models"

type Repository interface {
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	GetUsers() (*[]models.User, error)
	GetById(id interface{}) (*models.User, error)
	Login(u models.UserLogin) (models.User, error)
}
