package main

import (
	"context"
	"github.com/abolfazlalz/goasali/internal/users/db/migrate"
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/routers"
	"github.com/abolfazlalz/goasali/internal/users/services"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"github.com/abolfazlalz/goasali/pkg/config"
	"github.com/abolfazlalz/goasali/pkg/database"
	routes "github.com/abolfazlalz/goasali/pkg/http/routers"
	"github.com/abolfazlalz/goasali/pkg/multilingual"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

func main() {
	ctx := context.Background()

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

	var err error
	var databaseConfig *database.Database
	if databaseConfig, err = database.LoadDatabase(); err != nil {
		log.Fatalf("Error during loading database: %v", err)
	}
	var cacheInstance cache.Cache
	if cacheInstance, err = cache.New(ctx); err != nil {
		if err != nil {
			log.Fatalf("Error during loading cache: %v", err)
			return
		}
	}

	// Migrate modules models
	if err := databaseConfig.Migrate(migrate.UserMigration{}); err != nil {
		log.Fatalf("Database migration failed %v", err)
	}

	db := databaseConfig.DB
	//admin := services.NewAdmin(db, cacheInstance)
	role := services.NewRole(db, cacheInstance)
	err = role.CreatePermission(&models.Permission{Name: "Read"}, &models.Permission{Name: "Create"})
	if err != nil {
		log.Printf("Error during create permissions: %v", err)
	}

	router := routes.SetupRouter(databaseConfig.DB, m.Bundle, cacheInstance)
	router.AddRoutes(routers.NewUserRoute())

	if err := router.Listen(); err != nil {
		log.Fatalf("Error in listenning gin service server: %v", err)
	}
}
