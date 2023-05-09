package routes

import (
	"github.com/abolfazlalz/goasali/pkg/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

type IRouteCrud interface {
	Listen(prefix string)
	Create(*gin.Context)
	Get(*gin.Context)
	List(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type RouteCrud[T any] struct {
	IRouteCrud
	model      *T
	repository repositories.Interface[T]
}

func NewRouteCrud[T any](model *T, repository repositories.Interface[T]) RouteCrud[T] {
	return RouteCrud[T]{
		model:      model,
		repository: repository,
	}
}

func (rc RouteCrud[T]) Listen(prefix string, engine *gin.Engine) {
	grp := engine.Group(prefix)
	{
		grp.POST("/", rc.Create)
		// grp.PUT("/:id", rc.Update)
	}
}

func (rc RouteCrud[T]) Create(c *gin.Context) {
	if err := c.ShouldBind(rc.model); err != nil {
		c.JSON(400, "Can't bind model data")
		return
	}
	if rc.repository == nil {
		log.Fatal("repository not defined !")
		return
	}
	result := rc.repository.Create(rc.model)

	if result.Error != nil {
		c.JSON(500, result.Error)
		return
	}

	c.JSON(200, rc.model)
}
