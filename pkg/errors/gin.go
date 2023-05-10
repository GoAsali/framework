package errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

func (he *HttpError) i18nMessage(c *gin.Context, id string, field string) string {
	return he.mustLocalize(c, &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: id,
		},
		TemplateData: map[string]interface{}{
			"Field": field,
		},
	})
}

func (he *HttpError) handleValidationError(err validator.ValidationErrors, c *gin.Context) {
	message := he.mustLocalize(c, &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "FormValidationError",
		}})
	code := http.StatusBadRequest

	formErrors := map[string]string{}
	for _, field := range err {
		fmt.Println(field)
		id := "FormValidationErrorRequired"
		if field.Tag() == "required" {
			id = "FormValidationErrorRequired"
		} else if field.Tag() == "unique" {
			id = "FormValidationErrorUnique"
		}
		formErrors[field.Field()] = he.i18nMessage(c, id, field.Field())
	}

	c.JSON(code, gin.H{
		"message": message,
		"status":  false,
		"fields":  formErrors,
	})
}

func (he *HttpError) HandleGinError(err error, c *gin.Context) {
	message := err.Error()
	code := http.StatusInternalServerError

	fmt.Println(err)

	if verr, ok := err.(validator.ValidationErrors); ok {
		he.handleValidationError(verr, c)
		fmt.Printf("%v\n", verr[0].Tag())
		return
	}

	c.JSON(code, gin.H{
		"message": message,
		"status":  false,
	})
}
