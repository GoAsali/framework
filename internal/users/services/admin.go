package services

import (
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/db/repository"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"gorm.io/gorm"
)

type AdminInterface interface {
	UsersList(users *[]models.User) error
	CreateAdmin(user *models.User) error
}

type AdminService struct {
	AdminInterface
	repo *repository.UserRepository
	role *Role
	db   *gorm.DB
}

func NewAdmin(db *gorm.DB, cache cache.Cache) *AdminService {
	return &AdminService{
		repo: repository.NewUserRepository(db, cache),
		role: NewRole(db, cache),
	}
}

func (service *AdminService) UsersList(users *[]models.User) error {
	return service.repo.List(users).Error
}

func (service *AdminService) CreateAdmin(user *models.User) error {
	tx := service.repo.Create(user)
	return tx.Error
}

func (service *AdminService) CreateSuperAdmin(user *models.User) error {
	role := models.Role{Name: "Admin"}
	//permission
	return nil
}
