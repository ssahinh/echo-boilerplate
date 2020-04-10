package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	FullName string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);unique_index;not null"`
	Password string
	Token    string
}

func (u *User) Hash(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return hashedPassword, err
	}
	u.Password = string(hashedPassword)
	return hashedPassword, err
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
