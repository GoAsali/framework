package main

import (
	"github.com/abolfazlalz/goasali/internal/users/db/migrate"
	"github.com/abolfazlalz/goasali/internal/users/routers"
	"github.com/abolfazlalz/goasali/pkg/config"
	"github.com/abolfazlalz/goasali/pkg/database"
	routes "github.com/abolfazlalz/goasali/pkg/http/routers"
	"github.com/abolfazlalz/goasali/pkg/multilingual"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

func main() {
	log.Println("hello world")
	m := &multilingual.Multilingual{
		Bundle: i18n.NewBundle(language.English),
		Path:   "languages",
	}
	if err := m.Load(); err != nil {
		log.Fatalf("Error in loading languages files: %v", err)
	}

	if err := config.LoadEnvironments(); err != nil {
		log.Fatalf("Error in loading `.env` file: %v", err)
	}

	databaseConfig, err := database.LoadDatabase()

	// Migrate modules models
	if err := databaseConfig.Migrate(migrate.UserMigration{}); err != nil {
		log.Fatalf("Database migration failed %v", err)
	}

	if err != nil {
		log.Fatalf("Can't loading database: %v", err)
	}

	router := routes.SetupRouter(databaseConfig.DB, m.Bundle)
	router.AddRoutes(routers.NewUserRoute())

	if err := router.Listen(); err != nil {
		log.Fatalf("Error in listenning gin http server: %v", err)
	}
}
