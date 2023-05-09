package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       uint   `json:"id" gorm:"primarykey"`
	Username string `gorm:"unique" json:"username"`
	Password string
}

type UserRole struct {
	gorm.Model
	UserId uint
	RoleId uint
	Role   Role
	User   User
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
	return err
}

func (user *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
