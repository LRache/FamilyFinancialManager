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

type EmailConfig struct {
	SMTPHost  string
	SMTPPort  int
	Username  string
	Password  string
	FromEmail string
	Enabled   bool
}

var App *AppConfig
var Database *DatabaseConfig
var Email *EmailConfig

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

	Email = &EmailConfig{}
	Email.SMTPHost = v.GetString("email.smtp_host")
	Email.SMTPPort = v.GetInt("email.smtp_port")
	Email.Username = v.GetString("email.username")
	Email.Password = v.GetString("email.password")
	Email.FromEmail = v.GetString("email.from_email")
	Email.Enabled = v.GetBool("email.enabled")

	return nil
}
