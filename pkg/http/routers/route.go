package routes

import (
	"fmt"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"github.com/abolfazlalz/goasali/pkg/config"
	"github.com/abolfazlalz/goasali/pkg/http/controllers"
	"github.com/abolfazlalz/goasali/pkg/http/validations"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
	"log"
)

type Interface interface {
	Listen(*RouteModuleParams)
}

type RouteModule struct {
	Interface
	Route string
}

func NewRouteModule(route string) *RouteModule {
	return &RouteModule{Route: route}
}

type RouteModuleParams struct {
	Router *gin.RouterGroup
	DB     *gorm.DB
	*i18n.Bundle
	Cache      cache.Cache
	Controller *controllers.Controllers
}

type Route struct {
	*gin.Engine
	*i18n.Bundle
	DB        *gorm.DB
	appConfig *config.App
	cache     cache.Cache
}

func SetupRouter(db *gorm.DB, bundle *i18n.Bundle, cache cache.Cache) *Route {
	appConfig, err := config.LoadApp()
	if err != nil {
		log.Fatalf("Error during load app environments: %v", err)
	}
	if appConfig.Mode != "" {
		gin.SetMode(appConfig.Mode)
	}
	router := gin.Default()

	r := &Route{router, bundle, db, appConfig, cache}
	r.loadValidations()

	return r
}

func (r *Route) loadValidations() {
	if err := validations.AddDatabase(r.DB); err != nil {
		log.Fatalf("error during load database validation: %v", err)
	}
}

func (r *Route) AddRoutes(routes ...Interface) {
	r.Use(func(c *gin.Context) {
		c.Set("db", r.DB)
		c.Set("cache", r.cache)
		c.Set("bundle", r.Bundle)
	})
	for _, route := range routes {
		grp := r.Group("")
		ctrl := controllers.NewController(r.Bundle, r.cache, r.DB)
		route.Listen(&RouteModuleParams{grp, r.DB, r.Bundle, r.cache, ctrl})
	}
}

func (r *Route) Listen() error {
	addr := fmt.Sprintf("%s:%s", r.appConfig.Host, r.appConfig.Port)
	if err := r.Run(addr); err != nil {
		return err
	}

	r.Use(func(c *gin.Context) {
		c.Set("db", r.DB)
		c.Set("cache", r.cache)
		c.Set("bundle", r.Bundle)
	})

	return nil
}
