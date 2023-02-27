package infrastructure

import (
	"asalpolaki/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func AutoMigrate() string {
	err := DB.AutoMigrate(models.User{})
	if err != nil {
		return fmt.Sprintf("Something wrong happened in migrate db\n%s", err.Error())
	}

	return ""
}

// NewDatabase : initializes and returns mysql db
func NewDatabase() bool {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
		HOST, DBNAME)

	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")
		return false
	}

	DB = db
	AutoMigrate()

	fmt.Println("Database connection established")
	return true
}
