package handler

import (
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

func ResultToResponse[T any](ctx *gin.Context, result *service.Result[T], resp interface{}) {
	if result.IsFailed() {
		ctx.JSON(result.Code, gin.H{
			"message": result.Message,
		})
		return
	}
	ctx.JSON(result.Code, gin.H{
		"message":  result.Message,
		"response": resp, // 改为response以符合API文档
	})
}
