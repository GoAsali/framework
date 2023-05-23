package services

import (
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/db/repository"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"gorm.io/gorm"
)

type AdminInterface interface {
	UsersList(users *[]models.User) error
}

type AdminService struct {
	AdminInterface
	repo *repository.UserRepository
	db   *gorm.DB
}

func NewAdmin(db *gorm.DB, cache cache.Cache) *AdminService {
	return &AdminService{
		repo: repository.NewUserRepository(db, cache),
	}
}

func (service AdminService) UsersList(users *[]models.User) error {
	return service.repo.List(users).Error
}
