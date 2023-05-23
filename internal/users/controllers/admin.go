package controllers

import (
	"fmt"
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/services"
	"github.com/abolfazlalz/goasali/pkg/cache"
	"github.com/abolfazlalz/goasali/pkg/http/controllers"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

type AdminController struct {
	*controllers.Controllers
	service *services.AdminService
}

func NewAdmin(db *gorm.DB, bundle *i18n.Bundle, cache cache.Cache) *AdminController {
	return &AdminController{
		service:     services.NewAdmin(db, cache),
		Controllers: controllers.New(bundle, cache),
	}
}

func (ctrl AdminController) List(c *gin.Context) {
	users := &[]models.User{}
	if err := ctrl.service.UsersList(users); err != nil {
		ctrl.HandleHttp(c, ctrl.ErrorMessage(err.Error()))
	}

	fmt.Println(users)
	c.JSON(200, users)
}
