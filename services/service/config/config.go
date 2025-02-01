package config

import "github.com/spf13/viper"

var C Config

type (
	Config struct {
		AppConfig AppConfig
		Database  DatabaseConfig
		Jwt       JwtConfig
	}

	AppConfig struct {
		// ...
	}

	DatabaseConfig struct {
		ConnectionURL string
		Name          string
	}

	JwtConfig struct {
		Secret        string
		RefreshSecret string
		Expire        int
		RefreshExpire int
	}
)

func InitConfig() *Config {
	SetConfig(&C)
	return &C
}

func SetConfig(config *Config) {
	// Database
	config.Database.ConnectionURL = viper.GetString("database.url")
	config.Database.Name = viper.GetString("database.name")

	// JWT
	config.Jwt.Secret = viper.GetString("jwt.secret")
	config.Jwt.RefreshSecret = viper.GetString("jwt.refreshSecret")
	config.Jwt.Expire = viper.GetInt("jwt.expire")
	config.Jwt.RefreshExpire = viper.GetInt("jwt.refreshExpire")
}
