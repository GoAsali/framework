package routes

import (
	"github.com/gin-gonic/gin"
)

type Routes struct {
	router *gin.Engine
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewRoutes() Routes {
	r := Routes{
		router: gin.Default(),
	}

	r.router.Use(CORSMiddleware())

	v1 := r.router.Group("/v1")

	r.addUsersRoute(v1)

	r.router.Static("/uploads", "./public/uploads")

	return r
}

func (r Routes) Run(...string) error {
	return r.router.Run()
}
