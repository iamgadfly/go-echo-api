package usecase

import (
	"context"
	"github.com/AleksK1NG/api-mc/pkg/httpErrors"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/core"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/users/repository"
	"net/http"
)

type userUC struct {
	userRepo repository.UsersRepo
}

func NewUserUseCase(userRepo repository.UsersRepo) *userUC {
	return &userUC{
		userRepo: userRepo,
	}
}

func (u userUC) Register(ctx context.Context, user *core.User) (core.UserWithToken, error) {

	existsUser, err := u.userRepo.FindByEmail(ctx, user)
	if existsUser != nil || err == nil {
		return nil, httpErrors.NewRestErrorWithMessage(http.StatusBadRequest, httpErrors.ErrEmailAlreadyExists, nil)
	}
	//u.userRepo
}
