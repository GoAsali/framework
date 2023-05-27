package controllers

import (
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/services"
	"github.com/abolfazlalz/goasali/internal/users/types"
	"github.com/abolfazlalz/goasali/pkg/http/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleControllerI interface {
	CreateRole()
}

type RolesController struct {
	*controllers.Controllers
	service   services.RoleServiceI
	HttpError *types.UserHttpError
}

func NewRolesController(ctrl *controllers.Controllers) *RolesController {
	return &RolesController{
		Controllers: ctrl,
		service:     services.NewRole(ctrl.Db, ctrl.Cache),
		HttpError:   types.NewUserError(ctrl.HttpError),
	}
}

func (rc *RolesController) Create(c *gin.Context) {
	createRole := &CreateRole{}

	if err := c.ShouldBindJSON(createRole); err != nil {
		rc.HttpError.HandleGinError(err, c)
		return
	}

	role := models.Role{
		Name: createRole.Name,
	}

	rc.service.CreateRole(&role, createRole.PermissionsId)
	c.JSON(http.StatusCreated, gin.H{
		"message": rc.GetI18nMessage(c, "roles.created"),
	})
}
