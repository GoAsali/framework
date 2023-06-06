package cache

import (
	"context"
	"encoding/json"
	"github.com/abolfazlalz/goasali/pkg/config"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Redis struct {
	Cache
	client     *redis.Client
	ctx        context.Context
	defaultTTL time.Duration
}

type RedisOptionFunc func(*redisOptions)

type redisOptions struct {
	ctx context.Context
}

func defaultRedisOption() *redisOptions {
	return &redisOptions{
		ctx: context.Background(),
	}
}

func WithContext(ctx context.Context) RedisOptionFunc {
	return func(options *redisOptions) {
		options.ctx = ctx
	}
}

func NewRedis(optionsFunc ...RedisOptionFunc) *Redis {
	redisConfig, err := config.LoadCache()

	options := defaultRedisOption()
	for _, fn := range optionsFunc {
		fn(options)
	}

	if err != nil {
		log.Fatal("Error in load cache redisConfig: ", redisConfig)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Redis.Address,
		Password: redisConfig.Redis.Password,
		DB:       0,
	})
	return &Redis{client: rdb, ctx: options.ctx, defaultTTL: time.Hour}
}

func (r Redis) Set(item Item) error {
	data, err := json.Marshal(item.Value)
	if err != nil {
		return err
	}
	re := r.client.Set(r.ctx, item.Key, data, item.ttl())
	if err := re.Err(); err != nil {
		return err
	}
	return nil
}

func (r Redis) Get(key string, result interface{}) error {
	re, err := r.client.Get(r.ctx, key).Bytes()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		return err
	}
	log.Println(string(re))
	if err := json.Unmarshal(re, result); err != nil {
		return err
	}
	return nil
}

// Remember If a key has value return else, set value from function callback
func (r Redis) Remember(key string, ttl time.Duration, f func() interface{}) (interface{}, error) {
	var err error
	var result interface{}

	err = r.Get(key, &result)

	if err != nil {
		return "", err
	}

	// Still this key has value in cache
	if result != nil {
		return result, nil
	}

	result = f()

	if err := r.Set(Item{TTL: ttl, Key: key, Value: result}); err != nil {
		return "", err
	}

	return result, nil
}

func (r Redis) Forget(key ...string) error {
	return r.client.Del(r.ctx, key...).Err()
}
