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
	Tags           string `gorm:"not null"`
	Email          string `gorm:"not null;unique"`
	Phone          string
	City           string
	State          string
	FullAddress    string
	LastLogin      time.Time
	IsAdmin        bool
	IsSeller       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
