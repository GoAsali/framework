package services

import (
	"errors"
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/db/repository"
	"github.com/abolfazlalz/goasali/internal/users/utils/tokens"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"gorm.io/gorm"
)

type AuthServiceI interface {
	Login(user *models.User, username string, password string) (string, error)
	CreateAccount(user *models.User) (string, error)
	DeleteUser(user *models.User) error
}

var (
	UserUnauthorizedError = errors.New("username or password was incorrect")
)

type AuthService struct {
	AuthServiceI
	db    *gorm.DB
	repo  *repository.UserRepository
	token *tokens.Token
}

func NewAuthService(db *gorm.DB, cache cache.Cache) *AuthService {
	return &AuthService{
		db:    db,
		repo:  repository.NewUserRepository(db, cache),
		token: tokens.New(),
	}
}

func (as *AuthService) Login(user *models.User, username string, password string) (string, error) {
	if err := as.repo.FindByUsername(username, user); err != nil {
		return "", err
	}
	if user == nil || !user.CheckPasswordHash(password) {
		return "", UserUnauthorizedError
	}

	token, err := as.token.GenerateJwtToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (as *AuthService) CreateAccount(user *models.User) (string, error) {
	err := user.HashPassword()
	if err != nil {
		return "", err
	}
	tx := as.db.Create(user)

	if tx.Error != nil {
		return "", tx.Error
	}

	token, err := as.token.GenerateJwtToken(user)

	if err != nil {
		if user.Id != 0 {
			if err := as.DeleteUser(user); err != nil {
				return "", err
			}
		}
		return "", err
	}

	return token, nil
}

func (as *AuthService) DeleteUser(user *models.User) error {
	tx := as.db.Delete(user)
	as.db.Delete(&user, user.Id)
	return tx.Error
}
