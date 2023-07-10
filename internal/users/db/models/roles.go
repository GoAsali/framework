package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string
	RoleId      uint
	Role        *Role         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Permissions []*Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
	gorm.Model
	Name  string  `json:"name" gorm:"unique"`
	Roles []*Role `gorm:"many2many:role_permissions;"`
}
