package routes

import (
	"asalpolaki/controllers"
	"asalpolaki/middlewares"
	"asalpolaki/services"
	"github.com/gin-gonic/gin"
)

func (r Routes) addUsersRoute(rg *gin.RouterGroup) {
	router := rg.Group("/users")

	userController := controllers.UserController{JwtService: services.JWTService{}}

	router.POST("/auth/create", userController.CreateAccount)
	router.POST("/auth/login", userController.Login)
	router.GET("/auth", middlewares.AuthenticationMiddleware(), userController.UserDetails)
}
