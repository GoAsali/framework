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

func (UserRouter) listenAuth(route *routes.RouteModuleParams) {
	ctrl := controllers.NewAuthController(route.Controller)
	ctrl = controllers.NewAuthLogs(ctrl)
	grp := route.Router.Group("/auth")
	grp.POST("/register", ctrl.CreateAccount)
	grp.POST("/login", ctrl.Login)
	grp.POST("/refresh", ctrl.RefreshToken)
	grp.Use(middlewares.IsAuthMiddleware).GET("/", ctrl.Info)
}

func (UserRouter) listenRoles(route *routes.RouteModuleParams) {
	ctrl := controllers.NewRolesController(route.Controller)
	grp := route.Router.Group("/users/roles")
	grp.POST("/", ctrl.Create)
}

func (r UserRouter) Listen(route *routes.RouteModuleParams) {
	r.listenAuth(route)
	r.listenRoles(route)
}
