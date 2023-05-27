package services

import (
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/db/repository"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"gorm.io/gorm"
)

type RoleServiceI interface {
	CreateRole(role *models.Role, permissionsId []uint)
	RemoveRole(userId int)
}

type RoleService struct {
	RoleServiceI
	db   *gorm.DB
	repo repository.RoleRepositoryI
}

func NewRole(db *gorm.DB, cache cache.Cache) RoleServiceI {
	return &RoleService{
		db:   db,
		repo: repository.NewRole(db, cache),
	}
}
