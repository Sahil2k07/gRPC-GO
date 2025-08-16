package config

import (
	"os"
	"strings"
)

type serverConfig struct {
	ServerPort string   `toml:"server_port"`
	Origins    []string `toml:"origins"`
}

func loadProdServerConfig() serverConfig {
	org := os.Getenv("APP_ORIGINS")
	return serverConfig{
		ServerPort: os.Getenv("SERVER_PORT"),
		Origins:    strings.Split(org, ","),
	}
}

func GetServerConfig() serverConfig {
	return serverConfig{
		ServerPort: globalConfig.Server.ServerPort,
		Origins:    globalConfig.Server.Origins,
	}
}
