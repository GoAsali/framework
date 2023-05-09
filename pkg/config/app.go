package config

import (
	"fmt"
	"github.com/abolfazlalz/goasali/pkg/utils/slices"
	"github.com/caarlos0/env/v8"
	"github.com/gin-gonic/gin"
)

type AppMode string

const (
	ReleaseMode = gin.ReleaseMode
	DebugMode   = gin.DebugMode
	TestMode    = gin.TestMode
)

type App struct {
	Name string `env:"APP_NAME"`
	Host string `env:"APP_HOST"`
	Port string `env:"APP_PORT"`
	Mode string `env:"APP_MODE"`
}

func LoadApp() (*App, error) {
	appConfig := &App{}
	if err := env.Parse(appConfig); err != nil {
		return nil, err
	}
	if appConfig.Port == "" {
		appConfig.Port = "9000"
	}
	if appConfig.Mode == "" {
		appConfig.Mode = DebugMode
	}
	appModes := []string{ReleaseMode, DebugMode, TestMode}
	if !slices.Contains(appModes, appConfig.Mode) {
		return nil, fmt.Errorf("invalid app mode \"%s\"", appConfig.Mode)
	}
	return appConfig, nil
}
