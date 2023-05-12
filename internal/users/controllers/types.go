package controllers

type LoginUser struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type RegisterUser struct {
	Username        string `binding:"required,unique=users"`
	Password        string `binding:"required"`
	ConfirmPassword string `binding:"required" json:"confirm_password"`
	FirstName       string `binding:"required" json:"first_name"`
	LastName        string `binding:"required" json:"last_name"`
}
