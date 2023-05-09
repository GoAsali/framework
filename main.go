package main

import (
	"github.com/abolfazlalz/goasali/internal/users/db/migrate"
	"github.com/abolfazlalz/goasali/internal/users/routers"
	"github.com/abolfazlalz/goasali/pkg/config"
	"github.com/abolfazlalz/goasali/pkg/database"
	routes "github.com/abolfazlalz/goasali/pkg/http/routers"
	"github.com/abolfazlalz/goasali/pkg/languages"
	"log"
)

func main() {
	bundle := languages.LoadLanguages()

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

	router := routes.SetupRouter(databaseConfig.DB, bundle)
	router.AddRoutes(routers.NewUserRoute())

	if err := router.Listen(); err != nil {
		log.Fatalf("Error in listenning gin http server: %v", err)
	}
}
