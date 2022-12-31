package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID             string  `gorm:"primaryKey"`
	ProductName    string  `gorm:"not null"`
	ProductDetails string  `gorm:"not null"`
	Price          float64 `gorm:"not null"`
	Status         string  `gorm:"default: N"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
