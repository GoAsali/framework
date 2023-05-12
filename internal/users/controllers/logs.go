package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
)

type AuthLogsControllers struct {
	IAuthController
}

func NewAuthLogs(controller IAuthController) IAuthController {
	return &AuthLogsControllers{
		controller,
	}
}

func (logAuth *AuthLogsControllers) CreateAccount(c *gin.Context) {
	lastPrefix := log.Prefix()
	log.SetPrefix("[AuthLogsControllers][CreateAccount] ")

	log.Println("Start Creating account")

	defer func(prefix string) {
		log.Println("Create account method complete")
		log.SetPrefix(prefix)
	}(lastPrefix)

	logAuth.IAuthController.CreateAccount(c)
}

func (logAuth *AuthLogsControllers) Login(c *gin.Context) {
	lastPrefix := log.Prefix()
	log.SetPrefix("[AuthLogsControllers][Login] ")

	defer func(prefix string) {
		log.Println("login method complete")
		log.SetPrefix(prefix)
	}(lastPrefix)

	logAuth.IAuthController.Login(c)
}
