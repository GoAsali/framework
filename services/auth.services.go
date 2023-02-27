package services

import (
	"asalpolaki/infrastructure"
	"asalpolaki/models"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth(c *gin.Context, user *models.User) (err error) {
	bearerToken := c.Request.Header["Authorization"]
	token := strings.Split(bearerToken[0], " ")[1]
	service := JWTService{}
	data, err := service.EncodeJWT(token)
	if err != nil {
		return err
	}

	result := infrastructure.DB.First(user, "id=?", data["user_id"])
	if result.Error != nil {
		return result.Error
	}

	return nil
}
