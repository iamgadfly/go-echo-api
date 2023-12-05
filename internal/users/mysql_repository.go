package users

import "github.com/iamgadfly/go-echo-api/internal/models"

type Repository interface {
	Create(user *models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}
