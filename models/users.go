package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          int64  `gorm:"primaryKey;autoIncrement:true"`
	UUID        string `gorm:"not null;unique"`
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Email       string `gorm:"not null;unique"`
	Phone       string
	City        string
	State       string
	FullAddress string
	LastLogin   time.Time
	IsAdmin     bool `gorm:"not null;unique;default: 0"`
	IsSeller    bool `gorm:"not null;unique;default: 0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
