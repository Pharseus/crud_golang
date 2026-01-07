package entities

import "time"

type Product struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(100);not null"`
	SKU       string    `gorm:"type:varchar(50);unique;not null"`
	Price     float64   `gorm:"type:decimal(8,2);not null"`
	Stock     int       `gorm:"default:0"`
	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
