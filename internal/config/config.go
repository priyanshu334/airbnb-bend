package config

import "github.com/spf13/viper"

type Config struct {
	AppPort string
	DBUrl   string
}

func Load() *Config {
	viper.SetDefault("APP_PORT", 8000)
	viper.AutomaticEnv()

	return &Config{
		AppPort: viper.GetString("APP_PORT"),
		DBUrl:   viper.GetString("DB_URL"),
	}
}
