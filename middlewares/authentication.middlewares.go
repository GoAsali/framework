package middlewares

import (
	"asalpolaki/models"
	"asalpolaki/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header["Authorization"]
		if len(bearerToken) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "status": false})
			return
		}
		user := models.User{}
		err := services.Auth(c, &user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "status": false})
			return
		}
		c.Next()
	}
}
