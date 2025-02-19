package config

import "github.com/spf13/viper"

var C Config

type (
	Config struct {
		ServiceToken string
	}
)

func InitConfig() *Config {
	SetConfig(&C)
	return &C
}

func SetConfig(config *Config) {
	config.ServiceToken = viper.GetString("service.token")
}
