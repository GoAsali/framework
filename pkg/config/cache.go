package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
)

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     int    `env:"REDIS_PORT"`
	Address  string
	Password string `env:"REDIS_PASSWORD"`
}

type CacheConfig struct {
	Type  string `env:"CACHE_TYPE"`
	Redis *RedisConfig
}

func LoadRedis() (*RedisConfig, error) {
	redisConfig := &RedisConfig{}
	if err := env.Parse(redisConfig); err != nil {
		return nil, err
	}
	if redisConfig.Host == "" {
		redisConfig.Host = "localhost"
	}
	if redisConfig.Port == 0 {
		redisConfig.Port = 6379
	}

	redisConfig.Address = fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	return redisConfig, nil
}

func LoadCache() (*CacheConfig, error) {
	redis, err := LoadRedis()
	if err != nil {
		return nil, err
	}
	cache := &CacheConfig{
		Redis: redis,
	}
	if err := env.Parse(cache); err != nil {
		return nil, err
	}

	return cache, nil
}
