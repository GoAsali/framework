package tokens

import (
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/db/repository"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"gorm.io/gorm"
)

type UserJwt struct {
	token string
	db    *gorm.DB
	cache cache.Cache
}

func NewUserJwt(token string, db *gorm.DB, cache cache.Cache) *UserJwt {
	return &UserJwt{token, db, cache}
}

func (uj *UserJwt) User(user *models.User) error {
	token := New()
	claim, err := token.ValidateJwtToken(uj.token)
	if err != nil {
		return err
	}

	repo := repository.NewUserRepository(uj.db, uj.cache)

	return repo.FindByUsername(claim.Username, user)
}
