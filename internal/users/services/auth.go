package services

import (
	"errors"
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/db/repository"
	"github.com/abolfazlalz/goasali/internal/users/utils/tokens"
	"gorm.io/gorm"
)

var (
	UserUnauthorizedError = errors.New("username or password was incorrect")
)

type AuthService struct {
	db    *gorm.DB
	repo  *repository.UserRepository
	token *tokens.Token
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db,
		repository.NewUserRepository(db),
		tokens.New(),
	}
}

func (as *AuthService) Login(user *models.User, username string, password string) error {
	if err := as.repo.FindByUsername(username, user); err != nil {
		return err
	}
	if user != nil || user.CheckPasswordHash(password) {
		return UserUnauthorizedError
	}

	return nil
}

func (as *AuthService) CreateAccount() {}
