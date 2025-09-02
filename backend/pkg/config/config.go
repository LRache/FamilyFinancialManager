package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Host         string
	Port         string
	RouterPrefix string
	JwtSecret    string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
}

var App *AppConfig
var Database *DatabaseConfig

func Init() error {
	v := viper.New()

	v.SetConfigFile("config/config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	App = &AppConfig{}
	App.Host = v.GetString("app.host")
	App.Port = v.GetString("app.port")
	App.RouterPrefix = v.GetString("app.router_prefix")
	App.JwtSecret = v.GetString("app.jwt_secret")

	Database = &DatabaseConfig{}
	Database.Host = v.GetString("database.host")
	Database.Port = v.GetString("database.port")
	Database.Username = v.GetString("database.username")
	Database.Password = v.GetString("database.password")
	Database.Dbname = v.GetString("database.dbname")

	return nil
}
