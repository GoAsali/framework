package routers

import (
	"github.com/abolfazlalz/goasali/internal/users/controllers"
	"github.com/abolfazlalz/goasali/internal/users/middlewares"
	routes "github.com/abolfazlalz/goasali/pkg/http/routers"
)

type UserRouter struct {
	routes.Interface
}

func NewUserRoute() *UserRouter {
	return &UserRouter{routes.NewRouteModule("users")}
}

func (UserRouter) Listen(route *routes.RouteModuleParams) {
	ctrl := controllers.NewAuthController(route.DB, route.Bundle, route.Cache)
	ctrl = controllers.NewAuthLogs(ctrl)
	grp := route.Router.Group("/auth")
	grp.POST("/register", ctrl.CreateAccount)
	grp.POST("/login", ctrl.Login)
	grp.Use(middlewares.IsAuthMiddleware).GET("/", ctrl.Info)
}
