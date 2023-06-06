package cache

import (
	"context"
	"github.com/abolfazlalz/goasali/pkg/config"
	"log"
	"time"
)

type SupportValues interface {
	string | int
}

type Item struct {
	Key   string
	Value interface{}
	TTL   time.Duration
}

type Cache interface {
	Get(key string, value interface{}) error
	Set(item Item) error
	Remember(key string, ttl time.Duration, f func() interface{}) (interface{}, error)
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
		return NewRedis(WithContext(ctx)), nil
	}

	return nil, InvalidTypeError{Type: configCache.Type}
}

func (item *Item) ttl() time.Duration {
	ttl := time.Hour
	if item.TTL == 0 {
		return ttl
	}
	if item.TTL < time.Second {
		log.Printf("Too short ttl for key: %q: %s", item.Key, item.TTL)
		return ttl
	}
	return ttl
}
