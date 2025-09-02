package router

import (
	"backend/internal/handler"
	"backend/pkg/config"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	root := router.Group(config.Cfg.App.RouterPrefix)

	authGroup := root.Group("/auth")
	authGroup.POST("/register", handler.RegisterUser)

	return router
}
