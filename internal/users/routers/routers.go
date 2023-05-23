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

func (UserRouter) adminCtrl(route *routes.RouteModuleParams) {
	ctrl := controllers.NewAdmin(route.DB, route.Bundle, route.Cache)
	grp := route.Router.Group("/admin")
	grp.GET("/users", ctrl.List)
}

func (UserRouter) authCtrl(route *routes.RouteModuleParams) {
	ctrl := controllers.NewAuth(route.DB, route.Bundle, route.Cache)
	ctrl = controllers.NewAuthLogs(ctrl)

	grp := route.Router.Group("/auth")
	grp.POST("/register", ctrl.CreateAccount)
	grp.POST("/login", ctrl.Login)
	grp.POST("/refresh", ctrl.RefreshToken)
	grp.Use(middlewares.IsAuthMiddleware).GET("/", ctrl.Info)
}

func (ur UserRouter) Listen(route *routes.RouteModuleParams) {
	ur.authCtrl(route)
	ur.adminCtrl(route)
}
