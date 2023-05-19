package controllers

import (
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/services"
	"github.com/abolfazlalz/goasali/internal/users/types"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"github.com/abolfazlalz/goasali/pkg/http/controllers"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type IAuthController interface {
	Login(c *gin.Context)
	CreateAccount(c *gin.Context)
	Info(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type AuthController struct {
	IAuthController
	*controllers.Controllers
	HttpError   *types.UserHttpError
	authService services.AuthServiceI
}

func NewAuthController(db *gorm.DB, bundle *i18n.Bundle, cache cache.Cache) IAuthController {
	ctrl := controllers.NewController(bundle, cache)
	return &AuthController{
		Controllers: ctrl,
		authService: services.NewAuthServiceLogs(services.NewAuthService(db, cache)),
		HttpError:   types.NewUserError(ctrl.HttpError),
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
		"user":          user,
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"status":        true,
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
		"user":          user,
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"status":        true,
	})
}

func (ac *AuthController) RefreshToken(c *gin.Context) {
	body := RefreshTokenRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		ac.HttpError.HandleGinError(err, c)
		return
	}

	token, err := ac.authService.RefreshToken(body.Token)
	if err != nil {
		ac.HandleGinError(err, c)
		return
	}

	c.JSON(200, gin.H{
		"access_token": token,
		"status":       true,
	})
}

func (ac *AuthController) Info(c *gin.Context) {
	user, err := c.Get("user")
	if !err {
		ac.HttpError.HandleHttp(c, ac.HttpError.I18nErrorMessageConfig(c, "authorization.status_forbidden"), ac.HttpError.HttpCode(http.StatusForbidden))
		return
	}
	c.JSON(http.StatusOK, user)
}
