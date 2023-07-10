package controllers

type LoginUser struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type RegisterUser struct {
	Username        string `binding:"required,unique=service"`
	Password        string `binding:"required"`
	ConfirmPassword string `binding:"required" json:"confirm_password"`
	FirstName       string `binding:"required" json:"first_name"`
	LastName        string `binding:"required" json:"last_name"`
}

type RefreshTokenRequest struct {
	Token string `binding:"required" json:"refresh_token"`
}
