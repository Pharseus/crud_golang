package repository

import (
	"context"

	"github.com/Pharseus/crud_golang.git/api/entities"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order entities.Order) (entities.Order, error)
	FindOrder(ctx context.Context) ([]entities.Order, int64, error)
	FindOrderById(ctx context.Context, id int32) (entities.Order, error)
	UpdateOrderById(ctx context.Context, order entities.Order, id int32) (entities.Order, error)
	DeleteOrderById(ctx context.Context, id int32) error
}
