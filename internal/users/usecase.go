package users

import (
	"context"
	"github.com/AleksK1NG/api-mc/pkg/utils"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/core"
)

type UseCase interface {
	Register(ctx context.Context, user *core.User) (core.UserWithToken, error)
	Login(ctx context.Context, user *core.User) (core.UserWithToken, error)
	Update(ctx context.Context, user *core.User) (*core.User, error)
	FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*core.UsersList, error)
	GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*core.UsersList, error)
}
