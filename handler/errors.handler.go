package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown error"
}

func HandleError(c *gin.Context, err error) bool {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No data sent"})
		return true
	}
	out := make([]ErrorMsg, len(ve))
	for i, fe := range ve {
		out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
	return true
}
