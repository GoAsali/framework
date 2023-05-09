package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name   string
	RoleId uint
	Role   *Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Permission struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type RolePermission struct {
	gorm.Model
	RoleId       uint
	PermissionId uint
	Role         Role
	Permission   Permission
}
