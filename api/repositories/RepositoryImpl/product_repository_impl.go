package repositoryImpl

import (
	"context"

	"github.com/Pharseus/crud_golang.git/api/entities"
	repository "github.com/Pharseus/crud_golang.git/api/repositories/Repository"
	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) repository.ProductRepository {
	return &productRepositoryImpl{DB: db}
}

func (repo *productRepositoryImpl) CreateProduct(ctx context.Context, product entities.Product) (entities.Product, error) {
	result := repo.DB.WithContext(ctx).Create(&product)
	if result.Error != nil {
		return entities.Product{}, result.Error
	}
	return product, nil
}

func (repo *productRepositoryImpl) FindProduct(ctx context.Context) ([]entities.Product, int64, error) {
	var products []entities.Product
	var total int64

	if err := repo.DB.WithContext(ctx).Model(&entities.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	result := repo.DB.WithContext(ctx).Find(&products)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return products, total, nil
}

func (repo *productRepositoryImpl) FindProductById(ctx context.Context, id int32) (entities.Product, error) {
	var product entities.Product
	result := repo.DB.WithContext(ctx).First(&product, id)
	if result.Error != nil {
		return entities.Product{}, result.Error
	}
	return product, nil
}

func (repo *productRepositoryImpl) UpdateProductById(ctx context.Context, product entities.Product, id int32) (entities.Product, error) {
	result := repo.DB.WithContext(ctx).
		Model(&entities.Product{}).
		Where("id = ?", id).
		Updates(product)

	if result.Error != nil {
		return entities.Product{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entities.Product{}, gorm.ErrRecordNotFound
	}

	var updatedProduct entities.Product
	err := repo.DB.WithContext(ctx).First(&updatedProduct, id).Error
	if err != nil {
		return entities.Product{}, err
	}

	return updatedProduct, nil
}

func (repository *productRepositoryImpl) DeleteProductById(ctx context.Context, id int32) error {
	var product entities.Product
	result := repository.DB.WithContext(ctx).Model(&product).Where("id = ?", id).Delete(product)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// CreateProduct(ctx context.Context, product entities.Product) (entities.Product, error)
// func (repository *productRepositoryImpl) CreateProduct(ctx context.Context, product entities.Product) (entities.Product, error)
// func (repository *productRepositoryImpl) FindProduct(ctx context.Context) ([]entities.Product, int64, error)
// func (repository *productRepositoryImpl) FindProductById(ctx context.Context, id int32) (entities.Product, error)
// func (repository *productRepositoryImpl) UpdateProductById(ctx context.Context, product entities.Product, id int32) (entities.Product, error)
// func (repository *productRepositoryImpl) DeleteProductById(ctx context.Context, id int32) error
