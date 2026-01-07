package entities

import "time"

type Order struct {
	Id          int         `gorm:"primaryKey;autoIncrement"`
	UserId      int         `gorm:"not null"`
	User        User        `gorm:"foreignKey:UserId"`
	OrderNumber string      `gorm:"type:varchar(50);unique;not null"`
	Status      orderStatus `gorm:"type:enum('DRAFT','PAID','CANCELLED')"`
	TotalAmount float64     `gorm:"type:decimal(10,2);not null"`
	IsActive    bool        `gorm:"default:true"`
	CreatedAt   time.Time   `gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime"`
}

type orderStatus string

const (
	OrderStatusDraft     orderStatus = "DRAFT"
	OrderStatusPaid      orderStatus = "PAID"
	OrderStatusCancelled orderStatus = "CANCELLED"
)
