package services

import (
	"github.com/abolfazlalz/goasali/internal/users/db/repository"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
	*repository.UserRepository
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db:             db,
		UserRepository: repository.NewUserRepository(db),
	}
}

func (as AuthService) SignIn() {}

func (as AuthService) CreateAccount() {}
