package config

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
)

type appConfig struct {
	Database   databaseConfig `toml:"database"`
	JWT        jwtConfig      `toml:"jwt"`
	Origins    []string       `toml:"origins"`
	GrpcPort   string         `toml:"grpc_port"`
	ServerPort string         `toml:"server_port"`
	GrpcUrl    string         `toml:"grpc_url"`
	GrpcToken  string         `toml:"grpc_token"`
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
	origins := strings.Split(os.Getenv("APP_ORIGINS"), ",")

	globalConfig = appConfig{
		Database: databaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		JWT: jwtConfig{
			CookieName: os.Getenv("COOKIE_NAME"),
			Secret:     os.Getenv("JWT_SECRET"),
		},
		Origins:    origins,
		GrpcPort:   os.Getenv("GRPC_PORT"),
		ServerPort: os.Getenv("SERVER_PORT"),
		GrpcUrl:    os.Getenv("GRPC_URL"),
		GrpcToken:  os.Getenv("GRPC_TOKEN"),
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

		GenerateStockClient(globalConfig.GrpcUrl, globalConfig.GrpcToken)
	})

	return globalConfig
}
