package errors

import (
	routes "github.com/abolfazlalz/goasali/pkg/http/routers"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type I18nKeyString string

type HttpErrorMessage interface {
	I18nKeyString | string
}

type httpErrorConfig struct {
	httpCode int
	message  string
}

type OptionFunc func(*httpErrorConfig)

type HttpError struct {
	*i18n.Bundle
	*i18n.Localizer
}

func NewByContext(c *gin.Context) *HttpError {
	context := routes.NewContext(c)
	return &HttpError{
		Bundle: context.Bundle,
	}
}

func (he *HttpError) getLocalizer(c *gin.Context) *i18n.Localizer {
	accept := c.GetHeader("Accept-Language")
	return i18n.NewLocalizer(he.Bundle, accept)
}

func (he *HttpError) mustLocalize(c *gin.Context, lc *i18n.LocalizeConfig) string {
	return he.getLocalizer(c).MustLocalize(lc)
}

func defaultConfig() *httpErrorConfig {
	return &httpErrorConfig{
		httpCode: 500, message: "Unknown server error",
	}
}

func (he *HttpError) HttpCode(code int) OptionFunc {
	return func(opt *httpErrorConfig) {
		opt.httpCode = code
	}
}

func (he *HttpError) ErrorMessage(message string) OptionFunc {
	return func(opt *httpErrorConfig) {
		opt.message = message
	}
}

func (he *HttpError) I18nErrorMessageConfig(c *gin.Context, id string) OptionFunc {
	return func(opt *httpErrorConfig) {
		opt.message = he.I18nErrorMessage(c, id)
	}
}

func (he *HttpError) I18nErrorMessage(c *gin.Context, id string) string {
	return he.mustLocalize(c, &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: id,
		}})
}

func (he *HttpError) I18nErrorMessageOrDefault(c *gin.Context, id string, defValue string) string {
	value := he.I18nErrorMessage(c, id)
	if value == "" {
		return defValue
	}
	return value
}

func (he *HttpError) HandleHttp(c *gin.Context, configs ...OptionFunc) {
	config := defaultConfig()

	for _, cf := range configs {
		cf(config)
	}

	message := config.message

	response := gin.H{"message": message, "status": false}
	c.JSON(config.httpCode, response)
}
