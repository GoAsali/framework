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
			ID: "validation.error",
		}})
	code := http.StatusBadRequest

	formErrors := map[string]string{}
	for _, field := range err {
		id := "validation." + field.Tag()
		fieldKey := field.Field()
		keyId := fmt.Sprintf("validation.fields.%s", fieldKey)
		fieldName := he.I18nErrorMessageOrDefault(c, keyId, fieldKey)
		formErrors[fieldKey] = he.i18nMessage(c, id, fieldName)
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

	if verr, ok := err.(validator.ValidationErrors); ok {
		he.handleValidationError(verr, c)
		return
	}

	if verr, ok := err.(I18nMessageError); ok {
		he.HandleHttp(c, he.I18nErrorMessageConfig(verr.I18nId))
		return
	}

	he.HandleHttp(c, he.ErrorMessage(message), he.HttpCode(code))
}
