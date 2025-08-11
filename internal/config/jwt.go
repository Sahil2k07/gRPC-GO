package config

type jwtConfig struct {
	CookieName string `toml:"cookie_name"`
	Secret     string `toml:"secret"`
}

func GetJWTConfig() jwtConfig {
	return jwtConfig{
		CookieName: globalConfig.JWT.CookieName,
		Secret:     globalConfig.JWT.Secret,
	}
}
