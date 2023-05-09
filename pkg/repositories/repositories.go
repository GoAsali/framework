package repositories

import (
	"gorm.io/gorm"
)

type Interface[T any] interface {
	Create(model *T) (tx *gorm.DB)
	CreateMap(model map[string]string) (tx *gorm.DB)
	Get(id uint) *T
}

type Repository struct {
	Db *gorm.DB
}

func NewRepositoryInstance(db *gorm.DB) *Repository {
	return &Repository{
		Db: db,
	}
}
