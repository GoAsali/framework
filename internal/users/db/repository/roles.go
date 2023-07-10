package repository

import (
	"errors"
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"github.com/abolfazlalz/goasali/pkg/repositories"
	"gorm.io/gorm"
)

var (
	RoleNotFound = errors.New("role not found ")
)

type Role struct {
	*repositories.Repository[models.Role]
}

func NewRoleRepository(db *gorm.DB, cache cache.Cache) *Role {
	return &Role{
		Repository: repositories.NewRepositoryInstance[models.Role](db, cache),
	}
}

func (r *Role) AssignPermissions(roleId uint, permissions ...models.Permission) error {
	var role *models.Role
	r.Db.First(role, "id=?", roleId)
	if role == nil {
		return RoleNotFound
	}
	if err := r.Db.Where("id=?", roleId).Association("Permission").Append(permissions); err != nil {
		return err
	}
	return nil
}

func (r *Role) CreatePermission(permission *models.Permission) error {
	re := r.Db.Where("name=?", permission.Name).FirstOrCreate(permission)
	return re.Error
}

func (r *Role) DeletePermissionById(id ...uint) error {
	if len(id) < 1 {
		return errors.New("one id for delete a permission required")
	}
	r.Db.Delete(id)
	return nil
}
