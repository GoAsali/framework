package migrate

import (
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/pkg/database"
	"gorm.io/gorm"
)

type UserMigration struct {
	database.IMigrate
}

func (UserMigration) Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		models.User{},
		models.Role{},
		models.RolePermission{},
		models.UserRole{},
	)
	return err
}
