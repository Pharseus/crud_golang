package entities

import (
	"time"
)

type User struct {
	Id           int       `gorm:"primaryKey;autoIncrement"`
	Name         string    `gorm:"type:varchar(100);not null"`
	Email        string    `gorm:"type:varchar(100);unique;not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	IsActive     bool      `gorm:"default:true"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	// DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
