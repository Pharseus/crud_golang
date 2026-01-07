package repository

import (
	"context"

	"github.com/Pharseus/crud_golang.git/api/entities"
)

type PaymentRepository interface {
	PaymentOrder(ctx context.Context, payment entities.Payment) (entities.Payment, error)
	FindPayment(ctx context.Context) ([]entities.Payment, int64, error)
	FindPaymentById(ctx context.Context, id int32) (entities.Payment, error)
	UpdatePaymentById(ctx context.Context, payment entities.Payment, id int32) (entities.Payment, error)
	DeletePaymentById(ctx context.Context, id int32) error
}
