package repositories

import (
	"fmt"
	"math/rand"
)

type CacheConfigFn func(config *CacheConfig)

type Cache[T any] struct {
	*Repository[T]
	prefix string
	ttl    uint
}

type CacheConfig struct {
	prefix string
	ttl    uint
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func defaultConfig() *CacheConfig {
	return &CacheConfig{
		ttl:    1000,
		prefix: randStringRunes(10),
	}
}

// PrefixCacheRepoConfig Set cache id name prefix for cache repository
func PrefixCacheRepoConfig(prefix string) CacheConfigFn {
	return func(config *CacheConfig) {
		config.prefix = prefix
	}
}

// TTlCacheRepoConfig Set cache ttl for repository config
func TTlCacheRepoConfig(ttl uint) CacheConfigFn {
	return func(config *CacheConfig) {
		config.ttl = ttl
	}
}

// NewCache Make new cache repository support
// This service can save get repository service into cache
func NewCache[T any](repo *Repository[T], fns ...CacheConfigFn) *Cache[T] {
	config := defaultConfig()

	for _, fn := range fns {
		fn(config)
	}

	return &Cache[T]{
		repo,
		config.prefix,
		config.ttl,
	}
}

func (c Cache[T]) Get(id uint) *T {
	key := fmt.Sprintf("%s-get-%d", c.prefix, id)
	var result T
	_ = c.Cache.Remember(key, 100, &result, func(x interface{}) {
		x = c.Repository.Get(id)
	})

	return &result
}
