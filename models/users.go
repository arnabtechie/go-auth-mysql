package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Email       string `gorm:"not null;unique"`
	Phone       string
	City        string
	State       string
	FullAddress string
	LastLogin   time.Time
	IsAdmin     bool `gorm:"not null;default: 0"`
	IsSeller    bool `gorm:"not null;default: 0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
