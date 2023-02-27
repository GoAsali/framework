package repository

import (
	"asalpolaki/infrastructure"
	"asalpolaki/models"
)

func GetUserByUsername(user *models.User, username string) {
	infrastructure.DB.First(user, "username=?", username)
}

func GetUserById(user *models.User, id int) {
	infrastructure.DB.First(user, "id=?", id)
}

func CreateUser(user models.User) {
	infrastructure.DB.Create(user)
}
