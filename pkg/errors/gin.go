package errors

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

func (he *HttpError) handleValidationError(err validator.ValidationErrors, c *gin.Context) {
	message := he.mustLocalize(c, &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "FormValidationError",
		}})
	code := http.StatusBadRequest

	formErrors := map[string]string{}
	for _, field := range err {
		message := he.mustLocalize(c, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID: "FormValidationErrorRequired",
			},
			TemplateData: map[string]interface{}{
				"Field": field.Field(),
			},
		})
		formErrors[field.Field()] = message
	}

	c.JSON(code, gin.H{
		"message": message,
		"status":  false,
		"fields":  formErrors,
	})
}

func (he *HttpError) HandleGinError(err error, c *gin.Context) {
	message := ""
	code := http.StatusInternalServerError
	accept := c.GetHeader("Accept-Language")
	loc := i18n.NewLocalizer(he.Bundle, accept)

	if verr, ok := err.(validator.ValidationErrors); ok {
		he.handleValidationError(verr, c)
		return
	}

	c.JSON(code, gin.H{
		"message": message,
		"status":  false,
	})
}
