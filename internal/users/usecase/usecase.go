package usecase

import (
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/iamgadfly/go-echo-api/internal/users/repository"
	"github.com/iamgadfly/go-echo-api/pkg/jwt"
	"go.uber.org/zap"
)

type userUC struct {
	cfg      *config.Config
	userRepo repository.UsersRepo
	logger   *zap.SugaredLogger
}

func NewUserUseCase(cfg *config.Config, userRepo *repository.UsersRepo, logger *zap.SugaredLogger) userUC {
	return userUC{
		cfg:      cfg,
		userRepo: *userRepo,
		logger:   logger,
	}
}

func (u userUC) Register(user *models.User) (*models.User, error) {
	user.HashPassword()
	createUser, err := u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return createUser, nil
}

func (u userUC) GetUsers() ([]models.User, error) {
	userList, err := u.userRepo.GetUsers()
	if err != nil {
		return userList, err
	}

	return userList, nil
}

func (u userUC) Login(password, email string) (models.UserWithToken, error) {
	user, err := u.userRepo.Login(password, email)
	if err != nil {
		return models.UserWithToken{}, err
	}

	token, err := jwt.GenerateJWTToken(&user, u.cfg)
	if err != nil {
		return models.UserWithToken{}, err
	}

	return models.UserWithToken{
		User:  &user,
		Token: token,
	}, nil
}

func (u userUC) GetByID(id interface{}) (*models.User, error) {
	user, err := u.userRepo.GetById(id)
	u.logger.Info(id)
	if err != nil {
		return user, err
	}

	return user, nil
}
