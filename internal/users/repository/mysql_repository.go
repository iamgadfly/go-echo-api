package repository

import (
	"errors"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) FindByEmail(email string) (*models.User, error) {
	user := models.User{}
	r.db.Get(&user, FindByEmail, email)
	return &user, nil
}

func (r *UsersRepo) Create(user *models.User) (*models.User, error) {
	check, err := r.FindByEmail(user.Email)
	if err != nil {
		return check, err
	}
	if check.ID != 0 {
		return check, errors.New("user already was")
	}
	_, er := r.db.NamedExec(CreateUser, user)
	if er != nil {
		return user, er
	}
	res, _ := r.FindByEmail(user.Email)

	return res, nil
}

func (r *UsersRepo) GetUsers() ([]models.User, error) {
	var list []models.User
	err := r.db.Select(&list, GetUsers)
	if err != nil {
		return list, err
	}
	return list, nil
}

func (r *UsersRepo) Login(password, email string) (models.User, error) {
	user := models.User{}
	if err := r.db.Get(&user, FindByEmail, email); err != nil {
		return user, nil
	}

	if err := user.ComparePasswords(password); err != nil {
		return models.User{}, err
	}

	return user, nil
}
