package main

import (
	"backend/api/router"
	"backend/internal/middleware"
	"backend/pkg/config"

	"github.com/wonderivan/logger"
)

func main() {
	err := config.Init()
	if err != nil {
		logger.Error("Failed to init config.")
	}

	r := router.Init()

	r.Use(middleware.AuthVerify)

	err = r.Run(config.Cfg.App.Host + ":" + config.Cfg.App.Port)

	if err != nil {
		logger.Emer("Failed to start the server")
	}
}
