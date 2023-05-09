package controllers

import (
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/errors"
	"github.com/abolfazlalz/goasali/internal/users/services"
	"github.com/abolfazlalz/goasali/pkg/http/controllers"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

type AuthController struct {
	*controllers.Controllers
	HttpError   *errors.UserHttpError
	authService *services.AuthService
}

type LoginUser struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func NewAuthController(db *gorm.DB, bundle *i18n.Bundle) *AuthController {
	ctrl := controllers.NewController(bundle)
	return &AuthController{
		Controllers: ctrl,
		authService: services.NewAuthService(db),
		HttpError:   errors.NewUserError(ctrl.HttpError),
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	body := LoginUser{}
	if err := c.ShouldBindJSON(&body); err != nil {
		ac.HttpError.HandleGinError(err, c)
		return
	}

	user := &models.User{}
	if err := ac.authService.Login(user, body.Username, body.Password); err != nil {
		if err == services.UserUnauthorizedError {
			ac.HttpError.HandleHttp(c, ac.HttpCode(400), ac.ErrorMessage("Unauthorized"))

		}
		ac.HttpError.HandleHttp(c, ac.HttpCode(500), ac.ErrorMessage(err.Error()))
		return
	}

	c.JSON(200, body)
}

func (ac *AuthController) CreateAccount(c *gin.Context) {
	//ToDO: Handle it
	c.Abort()
}
