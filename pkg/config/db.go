package config

import (
	"github.com/caarlos0/env/v8"
)

type Database struct {
	Name     string `env:"DB_NAME"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Type     string `env:"DB_TYPE"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
}

func LoadDatabase() (*Database, error) {
	databaseConfig := &Database{}
	if err := env.Parse(databaseConfig); err != nil {
		return nil, err
	}
	return databaseConfig, nil
}
