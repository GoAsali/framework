package main

import (
	"context"
	"github.com/abolfazlalz/goasali/internal/users/db/migrate"
	"github.com/abolfazlalz/goasali/internal/users/routers"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"github.com/abolfazlalz/goasali/pkg/config"
	"github.com/abolfazlalz/goasali/pkg/database"
	routes "github.com/abolfazlalz/goasali/pkg/http/routers"
	"github.com/abolfazlalz/goasali/pkg/multilingual"
	"log"
)

func main() {
	ctx := context.Background()

	m := multilingual.New()

	if err := m.Load(); err != nil {
		log.Fatalf("Error in loading languages files: %v", err)
	}

	if err := config.LoadEnvironments(); err != nil {
		log.Fatalf("Error in loading `.env` file: %v", err)
	}

	databaseConfig, err := database.LoadDatabase()
	cache, err := cache.New(ctx)
	if err != nil {
		log.Fatalf("Error during loading cache: %v", err)
		return
	}

	// Migrate modules models
	if err := databaseConfig.Migrate(migrate.UserMigration{}); err != nil {
		log.Fatalf("Database migration failed %v", err)
	}

	if err != nil {
		log.Fatalf("Can't loading database: %v", err)
	}

	router := routes.SetupRouter(databaseConfig.DB, m.Bundle, cache)
	router.AddRoutes(routers.NewUserRoute())

	if err := router.Listen(); err != nil {
		log.Fatalf("Error in listenning gin http server: %v", err)
	}
}
