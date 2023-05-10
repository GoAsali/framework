package routes

import (
	"fmt"
	"github.com/abolfazlalz/goasali/pkg/config"
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
}

type Route struct {
	*gin.Engine
	DB *gorm.DB
	*i18n.Bundle
	appConfig *config.App
}

func SetupRouter(db *gorm.DB, bundle *i18n.Bundle) *Route {
	appConfig, err := config.LoadApp()
	if err != nil {
		log.Fatalf("Error during load app environments: %v", err)
	}
	if appConfig.Mode != "" {
		gin.SetMode(appConfig.Mode)
	}
	router := gin.Default()

	r := &Route{router, db, bundle, appConfig}
	r.loadValidations()

	return r
}

func (r *Route) loadValidations() {
	if err := validations.AddDatabase(r.DB); err != nil {
		log.Fatalf("error during load database validation: %v", err)
	}
}

func (r *Route) AddRoutes(routes ...Interface) {
	for _, route := range routes {
		grp := r.Group("")
		route.Listen(&RouteModuleParams{grp, r.DB, r.Bundle})
	}
}

func (r *Route) Listen() error {
	addr := fmt.Sprintf("%s:%s", r.appConfig.Host, r.appConfig.Port)
	return r.Run(addr)
}
