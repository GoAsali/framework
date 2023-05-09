package config

import (
	"github.com/abolfazlalz/goasali/pkg/utils/cli"
	"github.com/joho/godotenv"
)

// LoadEnvironments Load environments from file
func LoadEnvironments() error {
	file := ".env"
	if cli.HasArgsKey("env") {
		file = cli.GetArgsFromKey("env")
	}
	return LoadFile(file)
}

// LoadFile Load environments from custom file
func LoadFile(file string) error {
	return godotenv.Load(file)
}
