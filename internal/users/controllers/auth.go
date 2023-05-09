package controllers

import (
	"github.com/abolfazlalz/goasali/internal/users/services"
	"github.com/abolfazlalz/goasali/pkg/http/controllers"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

type AuthController struct {
	*controllers.Controllers
	*services.AuthService
}

type LoginUser struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func NewAuthController(db *gorm.DB, bundle *i18n.Bundle) *AuthController {
	return &AuthController{
		AuthService: services.NewAuthService(db),
		Controllers: controllers.NewController(bundle),
	}
}

func (ac AuthController) Login(c *gin.Context) {
	body := LoginUser{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		ac.HttpError.HandleGinError(err, c)
		return
	}
	c.JSON(200, body)
}
func (ac AuthController) CreateAccount(c *gin.Context) {
}
