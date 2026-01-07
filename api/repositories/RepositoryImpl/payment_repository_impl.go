package repositoryImpl

import (
	"context"

	"github.com/Pharseus/crud_golang.git/api/entities"
	repository "github.com/Pharseus/crud_golang.git/api/repositories/Repository"
	"gorm.io/gorm"
)

type paymentRepositoryImpl struct {
	DB *gorm.DB
}

func NewPaymentRepositoryImpl(db *gorm.DB) repository.PaymentRepository {
	return &paymentRepositoryImpl{DB: db}
}

// CreatePayment(ctx context.Context, Payment entities.Payment) (entities.Payment, error)
func (repository *paymentRepositoryImpl) PaymentOrder(ctx context.Context, payment entities.Payment) (entities.Payment, error) {
	result := repository.DB.WithContext(ctx).Create(&payment)
	if result.Error != nil {
		return entities.Payment{}, result.Error
	}
	return payment, nil
}
func (repository *paymentRepositoryImpl) FindPayment(ctx context.Context) ([]entities.Payment, int64, error) {
	var payment []entities.Payment
	var total int64

	if err := repository.DB.WithContext(ctx).Model(&entities.Payment{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	result := repository.DB.WithContext(ctx).Find(&payment)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return payment, total, nil
}
func (repository *paymentRepositoryImpl) FindPaymentById(ctx context.Context, id int32) (entities.Payment, error) {
	var payment entities.Payment
	result := repository.DB.WithContext(ctx).First(&payment, id)
	if result.Error != nil {
		return entities.Payment{}, result.Error
	}
	return payment, nil
}
func (repository *paymentRepositoryImpl) UpdatePaymentById(ctx context.Context, payment entities.Payment, id int32) (entities.Payment, error) {
	result := repository.DB.WithContext(ctx).
		Model(&entities.Payment{}).
		Where("id = ?", id).
		Updates(payment)

	if result.Error != nil {
		return entities.Payment{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entities.Payment{}, gorm.ErrRecordNotFound
	}

	var updatedOrder entities.Payment
	err := repository.DB.WithContext(ctx).First(&updatedOrder, id).Error
	if err != nil {
		return entities.Payment{}, err
	}

	return updatedOrder, nil
}
func (repository *paymentRepositoryImpl) DeletePaymentById(ctx context.Context, id int32) error {
	result := repository.DB.WithContext(ctx).Delete(&entities.Payment{}, id)
	return result.Error
}
