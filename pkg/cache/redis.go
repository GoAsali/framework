package cache

import (
	"context"
	"fmt"
	"github.com/abolfazlalz/goasali/pkg/config"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Redis struct {
	Cache
	client *redis.Client
	ctx    context.Context
}

func NewRedis(ctx context.Context) *Redis {
	redisConfig, err := config.LoadCache()

	if err != nil {
		log.Fatal("Error in load cache redisConfig: ", redisConfig)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Redis.Address,
		Password: redisConfig.Redis.Password,
		DB:       0,
	})
	return &Redis{client: rdb, ctx: ctx}
}

func (r Redis) Set(key string, value interface{}, ttl uint) error {
	du, err := time.ParseDuration(fmt.Sprintf("%ds", ttl))
	if err != nil {
		return err
	}
	re := r.client.Set(r.ctx, key, value, du)
	if err := re.Err(); err != nil {
		return err
	}
	return nil
}

func (r Redis) Get(key string) (interface{}, error) {
	re, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return re, nil
}

// Remember If a key has value return else, set value from function callback
func (r Redis) Remember(key string, ttl uint, f func() interface{}) (interface{}, error) {
	var err error
	var result interface{}

	result, err = r.Get(key)

	if err != nil {
		return "", err
	}

	// Still this key has value in cache
	if result != "" {
		return result, nil
	}

	result = f()

	if err := r.Set(key, result, ttl); err != nil {
		return "", err
	}

	return result, nil
}

func (r Redis) Forget(key ...string) error {
	return r.client.Del(r.ctx, key...).Err()
}
