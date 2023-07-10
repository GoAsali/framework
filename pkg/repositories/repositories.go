package repositories

import (
	"github.com/abolfazlalz/goasali/pkg/cache"
	"gorm.io/gorm"
)

type Interface[T any] interface {
	Create(model *T) (tx *gorm.DB)
	FirstOrCreate(model *T)
	CreateMap(model map[string]string) (tx *gorm.DB)
	Get(id uint) *T
	Update(id uint, model map[string]string) (tx *gorm.DB)
	Delete(id ...uint) (tx *gorm.DB)
	// List get list of dedicated models with a condition
	List(models *[]T, condition ...T) (tx *gorm.DB)
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

func (r Repository[T]) List(models *[]T, queryExecute ...ListQueryExecuteFn) (tx *gorm.DB) {
	query := defaultListQuery()
	for _, fn := range queryExecute {
		fn(query)
	}

	tx = r.Db.Limit(query.limit)
	tx = tx.Offset(query.offset)

	for _, condition := range query.conditions {
		tx = tx.Where(condition)
	}

	tx.Find(models)

	return
}

func (r Repository[T]) FirstOrCreate(model *T, conditions ...interface{}) error {
	if re := r.Db.FirstOrCreate(model, conditions); re.Error != nil {
		return re.Error
	}
	return nil
}
