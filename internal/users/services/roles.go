package services

import (
	"errors"
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/db/repository"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"gorm.io/gorm"
)

type RoleServiceI interface {
	CreateRole(*models.Role) error
	AssignRole(userId int, roleId int) error
}

type Role struct {
	RoleServiceI
	repo *repository.Role
}

func NewRole(db *gorm.DB, cache cache.Cache) *Role {
	return &Role{
		repo: repository.NewRoleRepository(db, cache),
	}
}

func (r *Role) CreateRole(role *models.Role) error {
	re := r.repo.Create(role)
	return re.Error
}

func (r *Role) CreateRoleIfNotExists(role *models.Role, permissions ...models.Permission) error {
	if re := r.repo.FirstOrCreate(role); re.Error() != "" {
		return errors.New(re.Error())
	}
	if err := r.repo.AssignPermissions(role.RoleId, permissions...); err != nil {
		return err
	}
	return nil
}

func (r *Role) AssignRole(userId int, roleId int) error {
	return nil
}

func (r *Role) CreatePermission(permissions ...*models.Permission) error {
	var id []uint
	var err error
	for _, permission := range permissions {
		err = r.repo.CreatePermission(permission)
		id = append(id, permission.ID)
	}

	if err != nil {
		r.repo.Delete(id...)
		return err
	}
	return nil
}
