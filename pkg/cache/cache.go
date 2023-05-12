package cache

import (
	"context"
	"github.com/abolfazlalz/goasali/pkg/config"
	"log"
)

type SupportValues interface {
	string | int
}

type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, ttl uint) error
	Remember(key string, ttl uint, f func() interface{}) (interface{}, error)
	Forget(key ...string) error
}

func New(ctx context.Context) (Cache, error) {
	prefix := log.Prefix()
	log.SetPrefix("[Cache Loading] - ")
	defer func(prefix string) {
		log.SetPrefix(prefix)
	}(prefix)
	log.Print("Start cache")

	configCache, err := config.LoadCache()
	if err != nil {
		return nil, err
	}
	if configCache.Type == "redis" {
		log.Print("Starting redis as cache")
		return NewRedis(ctx), nil
	}

	return nil, InvalidTypeError{Type: configCache.Type}
}
