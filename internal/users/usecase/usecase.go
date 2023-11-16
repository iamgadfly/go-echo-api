package usecase

import (
	"fmt"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/iamgadfly/go-echo-api/internal/users/repository"
)

type userUC struct {
	userRepo repository.UsersRepo
}

func NewUserUseCase(userRepo *repository.UsersRepo) userUC {
	return userUC{
		userRepo: *userRepo,
	}
}

// (core.UserWithToken, error)

func (u userUC) Register(ctx map[string]interface{}, user *models.User) (models.UserWithToken, error) {
	fmt.Println(ctx)
	return models.UserWithToken{}, nil

	//return models.UserWithToken{}, nil
	//existsUser, err := u.userRepo.FindByEmail(ctx, user)
	//if existsUser != nil || err == nil {
	//	return nil, httpErrors.NewRestErrorWithMessage(http.StatusBadRequest, httpErrors.ErrEmailAlreadyExists, nil)
	//}

	//return u.userRepo, nil
	//u.userRepo
}
