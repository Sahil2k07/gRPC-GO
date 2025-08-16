package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

type appConfig struct {
	Database databaseConfig `toml:"database"`
	JWT      jwtConfig      `toml:"jwt"`
	Grpc     grpcConfig     `toml:"grpc"`
	Server   serverConfig   `toml:"server"`
}

var (
	globalConfig appConfig
	once         sync.Once
)

func IsProduction() bool {
	env := os.Getenv("APP_ENV")
	return env == "PRODUCTION" || env == "STAGING"
}

func loadProdConfig() {
	globalConfig = appConfig{
		Database: loadDatabaseConfig(),
		JWT:      loadProdJwtConfig(),
		Grpc:     loadProdGrpcConfig(),
		Server:   loadProdServerConfig(),
	}
}

func loadDevConfig() {
	path, err := filepath.Abs("dev.toml")
	if err != nil {
		panic("failed to find config file path: " + err.Error())
	}

	if _, err := toml.DecodeFile(path, &globalConfig); err != nil {
		panic("failed to decode config file: " + err.Error())
	}
}

func LoadConfig() appConfig {
	once.Do(func() {
		if IsProduction() {
			loadProdConfig()
		} else {
			loadDevConfig()
		}

		GenerateStockClient(globalConfig.Grpc)
	})

	return globalConfig
}
