package repository

import (
	"context"
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

func (r *UsersRepo) FindByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	return user, nil
	//foundUser := &core.User{}
	//if err := r.db.QueryRowxContext(ctx, findUserByEmail, user.Email).StructScan(foundUser); err != nil {
	//	return nil, errors.Wrap(err, "authRepo.FindByEmail.QueryRowxContext")
	//}
	//return foundUser, nil
}
