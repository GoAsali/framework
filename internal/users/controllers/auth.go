package controllers

//TODO add jwt token for authorization

import (
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/errors"
	"github.com/abolfazlalz/goasali/internal/users/services"
	"github.com/abolfazlalz/goasali/pkg/http/controllers"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
	"log"
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

type RegisterUser struct {
	Username        string `binding:"required,unique=users"`
	Password        string `binding:"required"`
	ConfirmPassword string `binding:"required" json:"confirm_password"`
	FirstName       string `binding:"required" json:"first_name"`
	LastName        string `binding:"required" json:"last_name"`
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
	token, err := ac.authService.Login(user, body.Username, body.Password)
	if err != nil {
		if err == services.UserUnauthorizedError {
			ac.HttpError.HandleHttp(c, ac.HttpCode(400), ac.I18nErrorMessageConfig(c, "authorization.unauthorized"))
			return
		}
		ac.HttpError.HandleHttp(c, ac.HttpCode(500), ac.I18nErrorMessageConfig(c, "errors.internal_server"))
		return
	}

	c.JSON(200, gin.H{
		"user":   user,
		"token":  token,
		"status": true,
	})
}

func (ac *AuthController) CreateAccount(c *gin.Context) {
	body := RegisterUser{}

	if err := c.ShouldBindJSON(&body); err != nil {
		ac.HttpError.HandleGinError(err, c)
		return
	}

	if body.Password != body.ConfirmPassword {
		ac.HttpError.HandleHttp(c, ac.I18nErrorMessageConfig(c, "validation.password_same"))
		return
	}

	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Password:  body.Password,
		Username:  body.Username,
	}

	token, err := ac.authService.CreateAccount(&user)

	if err != nil {
		log.Printf("Error during create new user: %v", err)
		ac.HttpError.HandleHttp(c, ac.HttpCode(500), ac.I18nErrorMessageConfig(c, "errors.internal_server"))
		return
	}

	c.JSON(201, gin.H{
		"user":   user,
		"token":  token,
		"status": true,
	})
}
