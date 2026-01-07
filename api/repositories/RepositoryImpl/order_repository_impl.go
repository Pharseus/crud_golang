package repositoryImpl

import (
	"context"

	"github.com/Pharseus/crud_golang.git/api/entities"
	repository "github.com/Pharseus/crud_golang.git/api/repositories/Repository"
	"gorm.io/gorm"
)

type orderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepositoryImpl(db *gorm.DB) repository.OrderRepository {
	return &orderRepositoryImpl{DB: db}
}

// CreateOrder(ctx context.Context, Order entities.Order) (entities.Order, error)
func (repository *orderRepositoryImpl) CreateOrder(ctx context.Context, order entities.Order) (entities.Order, error) {
	result := repository.DB.WithContext(ctx).Create(&order)
	if result.Error != nil {
		return entities.Order{}, result.Error
	}
	return order, nil
}
func (repository *orderRepositoryImpl) FindOrder(ctx context.Context) ([]entities.Order, int64, error) {

	var order []entities.Order
	var total int64

	if err := repository.DB.WithContext(ctx).Model(&entities.Order{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	result := repository.DB.WithContext(ctx).Find(&order)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return order, total, nil
}
func (repository *orderRepositoryImpl) FindOrderById(ctx context.Context, id int32) (entities.Order, error) {
	var order entities.Order
	result := repository.DB.WithContext(ctx).First(&order, id)
	if result.Error != nil {
		return entities.Order{}, result.Error
	}
	return order, nil
}
func (repository *orderRepositoryImpl) UpdateOrderById(ctx context.Context, order entities.Order, id int32) (entities.Order, error) {
	result := repository.DB.WithContext(ctx).
		Model(&entities.Order{}).
		Where("id = ?", id).
		Updates(order)

	if result.Error != nil {
		return entities.Order{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entities.Order{}, gorm.ErrRecordNotFound
	}

	var updatedOrder entities.Order
	err := repository.DB.WithContext(ctx).First(&updatedOrder, id).Error
	if err != nil {
		return entities.Order{}, err
	}

	return updatedOrder, nil
}
func (repository *orderRepositoryImpl) DeleteOrderById(ctx context.Context, id int32) error {
	result := repository.DB.WithContext(ctx).Delete(&entities.Order{}, id)
	return result.Error
}
