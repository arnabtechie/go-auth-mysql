package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	FirstName   string `gorm:"not null" json:"first_name" validate:"required,gte=1"`
	LastName    string `gorm:"not null" json:"last_name" validate:"required,gte=1"`
	Password    string `gorm:"not null" json:"password" validate:"required,gte=4"`
	Email       string `gorm:"unique" json:"email" validate:"required,email"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	FullAddress string `json:"full_address,omitempty"`
	IsAdmin     bool   `gorm:"not null;default: 0"`
	IsSeller    bool   `gorm:"not null;default: 0"`
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
