package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id           uint   `json:"id" gorm:"primarykey"`
	Username     string `gorm:"unique" json:"username"`
	Password     string
	passwordHash bool
}

type UserRole struct {
	gorm.Model
	UserId uint
	RoleId uint
	Role   Role
	User   User
}

func (user *User) HashPassword() error {
	if user.passwordHash {
		return nil
	}
	user.passwordHash = true
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
	return err
}

// Hooks
//
//func (user *User) BeforeCreate(*gorm.DB) error {
//	return user.HashPassword()
//}
//
//func (user *User) BeforeSave(*gorm.DB) error {
//	return user.HashPassword()
//}
