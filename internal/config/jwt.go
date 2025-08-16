package config

import "os"

type jwtConfig struct {
	CookieName string `toml:"cookie_name"`
	Secret     string `toml:"secret"`
}

func loadProdJwtConfig() jwtConfig {
	return jwtConfig{
		CookieName: os.Getenv("COOKIE_NAME"),
		Secret:     os.Getenv("JWT_SECRET"),
	}
}

func GetJWTConfig() jwtConfig {
	return jwtConfig{
		CookieName: globalConfig.JWT.CookieName,
		Secret:     globalConfig.JWT.Secret,
	}
}
