package repository

import (
	"context"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/core"
	"github.com/jmoiron/sqlx"
)

type UsersRepo struct {
	db *sqlx.DB
}

func (UsersRepo) NewUsersRepository(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) FindByEmail(ctx context.Context, user *core.User) (*core.User, error) {
	return user, nil
	//foundUser := &core.User{}
	//if err := r.db.QueryRowxContext(ctx, findUserByEmail, user.Email).StructScan(foundUser); err != nil {
	//	return nil, errors.Wrap(err, "authRepo.FindByEmail.QueryRowxContext")
	//}
	//return foundUser, nil
}
