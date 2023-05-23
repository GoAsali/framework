package config

import (
	"github.com/caarlos0/env/v8"
	"time"
)

type TokenConfig struct {
	AccessLifeTime  time.Duration `env:"TOKEN_ACCESS_LIFETIME"`
	RefreshLifeTime time.Duration `env:"TOKEN_REFRESH_TIMELIFE"`
}

func LoadTokenConfig() (tokenConfig *TokenConfig, err error) {
	tokenConfig = &TokenConfig{}
	err = env.Parse(tokenConfig)

	if tokenConfig.RefreshLifeTime == 0 {
		tokenConfig.RefreshLifeTime, err = time.ParseDuration("24h")
		if err == nil {
			tokenConfig.RefreshLifeTime *= 30
		}
	}

	if tokenConfig.AccessLifeTime == 0 {
		tokenConfig.AccessLifeTime, err = time.ParseDuration("30m")
	}

	return
}
