package database

import (
	"fmt"
	"github.com/abolfazlalz/goasali/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	*gorm.DB
}

type Config struct {
	*config.Database
}

func loadConfig() *Config {
	if dbConfig, err := config.LoadDatabase(); err != nil {
		log.Fatalf("Error during parse enviroments for database config: %v", err)
	} else {
		return &Config{
			Database: dbConfig,
		}
	}

	return nil
}

func LoadDatabase() (*Database, error) {
	defer func(prefix string) {
		log.SetPrefix("")
	}(log.Prefix())

	log.SetPrefix("database: ")
	log.Println("Connect to db...")

	var db *gorm.DB
	var err error

	dbConfig := loadConfig()
	if db, err = dbConfig.loadDatabase(); err != nil {
		return nil, err
	}

	log.Printf("Successfully connected to %s database", dbConfig.Type)

	return &Database{DB: db}, nil
}

func (c *Config) mysql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.Name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func (c *Config) sqlite() (*gorm.DB, error) {
	fileName := fmt.Sprintf("%s.sqlite", c.Name)
	return gorm.Open(sqlite.Open(fileName))
}

func (c *Config) loadDatabase() (*gorm.DB, error) {
	switch c.Type {
	case "mysql":
		return c.mysql()
	case "sqlite":
		return c.sqlite()
	}

	return nil, fmt.Errorf("unknown data type, db-type: %s", c.Type)
}
