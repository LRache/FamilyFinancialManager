package config

import (
	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Host         string
	Port         string
	RouterPrefix string
	JwtSecret    string
}

func Init() error {
	v := viper.New()

	v.SetConfigFile("config/config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	Cfg = &Config{}
	Cfg.App.Host = v.GetString("app.host")
	Cfg.App.Port = v.GetString("app.port")
	Cfg.App.RouterPrefix = v.GetString("app.router_prefix")
	Cfg.App.JwtSecret = v.GetString("app.jwt_secret")

	return nil
}
