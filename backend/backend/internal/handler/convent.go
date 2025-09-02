package handler

import (
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

func ResultToResponse(ctx *gin.Context, result *service.Result) {
	ctx.JSON(result.Code, gin.H{
		"message": result.Message,
		"resp":    result.Data,
	})
}
