package main

import (
	"backend/api/router"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

func main() {
	err := config.Init()
	if err != nil {
		logger.Error("Failed to init config.")
	}

	err = repository.Init()
	if err != nil {
		logger.Emer("Failed to init database: ", err.Error())
	}

	r := gin.Default()

	r.Use(middleware.Cors)

	router.Init(r)

	err = r.Run(config.App.Host + ":" + config.App.Port)

	if err != nil {
		logger.Error("Failed to start the server")
	}
}
