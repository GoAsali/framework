package controllers

import (
	"github.com/abolfazlalz/goasali/pkg/cache"
	"github.com/abolfazlalz/goasali/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Controllers struct {
	*i18n.Bundle
	*errors.HttpError
	Cache cache.Cache
}

func New(bundle *i18n.Bundle, cache cache.Cache) *Controllers {
	return &Controllers{
		Bundle:    bundle,
		HttpError: &errors.HttpError{Bundle: bundle},
		Cache:     cache,
	}
}

func (ctrl Controllers) LoadLocalize(c *gin.Context) *i18n.Localizer {
	accept := c.GetHeader("Accept-Language")
	return i18n.NewLocalizer(ctrl.Bundle, accept)
}
