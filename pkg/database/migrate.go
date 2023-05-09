package database

import "gorm.io/gorm"

type IMigrate interface {
	Migrate(*gorm.DB) error
}

func (d Database) Migrate(migrations ...IMigrate) error {
	for _, migrate := range migrations {
		if err := migrate.Migrate(d.DB); err != nil {
			return err
		}
	}

	return nil
}
