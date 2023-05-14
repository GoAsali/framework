package repository

import (
	"errors"
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"github.com/abolfazlalz/goasali/pkg/repositories"

	"gorm.io/gorm"
)

var (
	UsernameNotFound = errors.New("username not found")
)

type UserRepository struct {
	repositories.Repository[models.User]
}

func NewUserRepository(db *gorm.DB, cache cache.Cache) *UserRepository {
	return &UserRepository{
		Repository: *repositories.NewRepositoryInstance[models.User](db, cache),
	}
}

func (ur *UserRepository) Create(user *models.User) (tx *gorm.DB) {
	{
		found := ur.CheckUsernameExists(user.Username)

		if found {
			return &gorm.DB{
				Error: UserAlreadyExistsError,
			}
		}
	}
	if err := user.HashPassword(); err != nil {
		return &gorm.DB{
			Error: HashPasswordError{err},
		}
	}
	return ur.Db.Create(user)
}

func (ur *UserRepository) Get(id uint) *models.User {
	user := models.User{Id: id}
	ur.Db.Where(user).First(&user)
	return &user
}

func (ur *UserRepository) CreateMap(model map[string]string) *gorm.DB {
	user := models.User{
		Username: model["Username"],
		Password: model["Password"],
	}
	return ur.Create(&user)
}

func (ur *UserRepository) FindByUsername(username string, user *models.User) error {
	result := ur.Db.Where(&models.User{Username: username}).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return result.Error
}

func (ur *UserRepository) CheckUsernameExists(username string) bool {
	var user *models.User
	if err := ur.FindByUsername(username, user); err != nil {
		return false
	}
	return user != nil
}

func (ur *UserRepository) AddRole() {}
