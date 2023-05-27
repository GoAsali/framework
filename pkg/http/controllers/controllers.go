package controllers

import (
	"github.com/abolfazlalz/goasali/pkg/cache"
	"github.com/abolfazlalz/goasali/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

type Controllers struct {
	*i18n.Bundle
	*errors.HttpError
	Cache cache.Cache
	Db    *gorm.DB
}

func NewController(bundle *i18n.Bundle, cache cache.Cache, db *gorm.DB) *Controllers {
	return &Controllers{
		Bundle:    bundle,
		HttpError: &errors.HttpError{Bundle: bundle},
		Db:        db,
		Cache:     cache,
	}
}

func (ctrl Controllers) GetI18nMessage(c *gin.Context, messageId string) string {
	accept := c.GetHeader("Accept-Language")
	loc := i18n.NewLocalizer(ctrl.Bundle, accept)
	return loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: messageId,
		}})
}

func (ctrl Controllers) LoadLocalize(c *gin.Context) *i18n.Localizer {
	accept := c.GetHeader("Accept-Language")
	return i18n.NewLocalizer(ctrl.Bundle, accept)
}
