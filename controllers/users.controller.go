package controllers

import (
	"asalpolaki/handler"
	"asalpolaki/infrastructure"
	"asalpolaki/models"
	"asalpolaki/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	JwtService services.JWTService
}

type CreateAccountType struct {
	Username  string `form:"username" json:"username"  binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
}

type LoginType struct {
	Username string `form:"username" json:"username" binding:"required"`
	Pass     string `form:"password" json:"password" binding:"required"`
}

func (controller UserController) CheckUsernameEmail(username string, email string) bool {
	var users []models.User
	infrastructure.DB.Find(&users, "email=? or username=?", email, username)
	if len(users) > 0 {
		return true
	}
	return false
}

func (controller UserController) CreateAccount(c *gin.Context) {
	data := &CreateAccountType{}

	err := c.Bind(data)
	if err != nil {
		handler.HandleError(c, err)
		return
	}

	if controller.CheckUsernameEmail(data.Username, data.Email) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Email or username already taken"})
		return
	}

	password := controller.JwtService.HashPassword(data.Password)
	user := models.User{Username: data.Username, Lastname: data.LastName, Name: data.FirstName, Email: data.Email, Password: password}
	result := infrastructure.DB.Create(&user)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error in create user", "error": result.Error})
		return
	}

	message := fmt.Sprintf("Welcome %s", data.Username)
	token, _ := controller.JwtService.GenerateJWT(user)
	c.JSON(http.StatusCreated, gin.H{"message": message, "user": user, "token": token})
}

func (controller UserController) Login(c *gin.Context) {
	data := &LoginType{}
	err := c.Bind(data)

	if err != nil {
		handler.HandleError(c, err)
		return
	}

	password := data.Pass
	username := data.Username

	var user models.User
	result := infrastructure.DB.First(&user, "`username`=? or `email`=?", username, username)

	if result.RowsAffected == 0 || !controller.JwtService.CheckPasswordHash(password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username or Password was incorrect"})
		return
	}

	message := fmt.Sprintf("Welcome %s", data.Username)
	token, _ := controller.JwtService.GenerateJWT(user)
	c.JSON(http.StatusOK, gin.H{"message": message, "user": user, "token": token})
}

func (controller UserController) UserDetails(c *gin.Context) {
	var user models.User
	err := services.Auth(c, &user)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Your information", "user": user})
}
