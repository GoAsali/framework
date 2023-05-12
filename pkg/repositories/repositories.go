package repositories

import (
	"github.com/abolfazlalz/goasali/pkg/cache"
	"gorm.io/gorm"
)

type Interface[T any] interface {
	Create(model *T) (tx *gorm.DB)
	CreateMap(model map[string]string) (tx *gorm.DB)
	Get(id uint) *T
}

type Repository[T any] struct {
	Interface[T]
	Cache cache.Cache
	Db    *gorm.DB
}

func NewRepositoryInstance[T any](db *gorm.DB, cache cache.Cache) *Repository[T] {
	return &Repository[T]{
		Db:    db,
		Cache: cache,
	}
}
