package repository

import (
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"github.com/abolfazlalz/goasali/pkg/repositories"
	"gorm.io/gorm"
	"log"
)

type RoleRepositoryI interface {
	CreatePermission(permission models.Permission)
}

type RoleRepository struct {
	RoleRepositoryI
	repositories.Repository[models.Role]
}

func NewRole(db *gorm.DB, cache cache.Cache) RoleRepositoryI {
	return &RoleRepository{
		Repository: *repositories.NewRepositoryInstance[models.Role](db, cache),
	}
}

func (rr RoleRepository) Create(role *models.Role) (tx *gorm.DB) {
	tx = rr.Db.Create(role)
	return
}

func (rr RoleRepository) Get(id uint) (role *models.Role) {
	role = &models.Role{}
	tx := rr.Db.Where("id=?", id).First(role)
	if err := tx.Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("Error during get role by id: %v", tx.Error)
	}
	return
}
