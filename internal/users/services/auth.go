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
	Login(user *models.User, username string, password string) (*Token, error)
	CreateAccount(user *models.User) (*Token, error)
	DeleteUser(user *models.User) error
	RefreshToken(refreshToken string) (string, error)
}

type Token struct {
	RefreshToken string
	AccessToken  string
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

func (as *AuthService) getTokens(user *models.User) (token *Token, err error) {
	token = &Token{}
	token.RefreshToken, err = as.token.GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}
	token.AccessToken, err = as.token.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}
	return
}

func (as *AuthService) Login(user *models.User, username string, password string) (*Token, error) {
	if err := as.repo.FindByUsername(username, user); err != nil {
		return nil, err
	}
	if user == nil || !user.CheckPasswordHash(password) {
		return nil, UserUnauthorizedError
	}

	token, err := as.getTokens(user)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (as *AuthService) CreateAccount(user *models.User) (*Token, error) {
	err := user.HashPassword()
	if err != nil {
		return nil, err
	}
	tx := as.db.Create(user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	token, err := as.getTokens(user)

	if err != nil {
		if user.Id != 0 {
			if err := as.DeleteUser(user); err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	return token, nil
}

func (as *AuthService) DeleteUser(user *models.User) error {
	tx := as.db.Delete(user)
	as.db.Delete(&user, user.Id)
	return tx.Error
}

func (as *AuthService) RefreshToken(refreshToken string) (string, error) {
	claim, err := as.token.ValidateJwtToken(refreshToken)
	if err != nil {
		return "", err
	}

	if claim.Username == "" {
		return "", tokens.NewExpireTokenError()
	}

	user := &models.User{}
	if err := as.repo.FindByUsername(claim.Username, user); err != nil {
		return "", err
	}

	if user == nil {
		return "", repository.UsernameNotFound
	}

	accessToken, err := as.token.GenerateAccessToken(user)
	return accessToken, err
}
