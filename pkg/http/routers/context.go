package routes

import (
	"github.com/abolfazlalz/goasali/pkg/cache"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
	"log"
)

type Context struct {
	*gin.Context
	*gorm.DB
	cache.Cache
	*i18n.Bundle
}

func NewContext(c *gin.Context) *Context {
	context := &Context{}
	if db, exists := c.Get("db"); exists {
		context.DB = db.(*gorm.DB)
	} else {
		log.Println("Db not found in gin context, please check route configuration")
	}
	if bundle, exists := c.Get("bundle"); exists {
		context.Bundle = bundle.(*i18n.Bundle)
	} else {
		log.Println("Bundle not found in gin context, please check route configuration")
	}
	if c, exists := c.Get("cache"); exists {
		context.Cache = c.(cache.Cache)
	} else {
		log.Println("Cache not found in gin context, please check route configuration")
	}
	return context
}
