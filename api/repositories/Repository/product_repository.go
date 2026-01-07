package repository

import (
	"context"

	"github.com/Pharseus/crud_golang.git/api/entities"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product entities.Product) (entities.Product, error)
	FindProduct(ctx context.Context) ([]entities.Product, int64, error)
	FindProductById(ctx context.Context, id int32) (entities.Product, error)
	UpdateProductById(ctx context.Context, product entities.Product, id int32) (entities.Product, error)
	DeleteProductById(ctx context.Context, id int32) error
}
