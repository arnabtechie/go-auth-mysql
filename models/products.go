package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID             int64  `gorm:"primaryKey;autoIncrement:true"`
	UUID           string `gorm:"not null;unique"`
	ProductName    string `gorm:"not null"`
	ProductDetails string `gorm:"not null"`
	Status         string `gorm:"default: N"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
