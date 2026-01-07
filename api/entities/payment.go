package entities

import "time"

type Payment struct {
	Id            int           `gorm:"primaryKey;autoIncrement"`
	OrderId       int           `gorm:"not null"`
	Order         Order         `gorm:"foreignKey:OrderId"`
	PaymentNumber string        `gorm:"type:varchar(50);unique;not null"`
	Method        paymentMethod `gorm:"type:enum('VA','CARD','CASH')"`
	Amount        float64       `gorm:"type:decimal(10,2);not null"`
	Status        paymentStatus `gorm:"type:enum('PENDING','SUCCESS','FAILED')"`
	PaidAt        *time.Time    `gorm:"default:null"`
	IsActive      bool          `gorm:"default:true"`
	CreatedAt     time.Time     `gorm:"autoCreateTime"`
	UpdatedAt     time.Time     `gorm:"autoUpdateTime"`
}

type paymentMethod string

const (
	PaymentMethodVA   paymentMethod = "VA"
	PaymentMethodCard paymentMethod = "CARD"
	PaymentMethodCash paymentMethod = "CASH"
)

type paymentStatus string

const (
	PaymentStatusPending paymentStatus = "PENDING"
	PaymentStatusSuccess paymentStatus = "SUCCESS"
	PaymentStatusFailed  paymentStatus = "FAILED"
)
