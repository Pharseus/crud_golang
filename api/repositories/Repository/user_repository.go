package repository

import (
	"context"

	"github.com/Pharseus/crud_golang.git/api/entities"
)

type Pagination struct {
	Limit  int
	Offset int
}

type UserRepository interface {
	CreateUser(ctx context.Context, user entities.User) (entities.User, error)
	FindUser(ctx context.Context, pagination Pagination) ([]entities.User, int64, error)
	FindUserById(ctx context.Context, id int32) (entities.User, error)
	UpdateUserById(ctx context.Context, user entities.User, id int32) (entities.User, error)
	DeleteUserById(ctx context.Context, id int32) error
}
