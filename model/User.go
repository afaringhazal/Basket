package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `validate:"required" gorm:"unique"`
	Password string `validate:"required"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = string(bytes)
	return err
}

func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
