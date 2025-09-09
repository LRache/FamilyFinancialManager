package handler

import (
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

// DebugRequestBody 调试函数：查看请求体数据
func DebugRequestBody(ctx *gin.Context) {
	// 读取原始请求体
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logger.Warn("Failed to read request body:", err.Error())
		return
	}

	// 记录原始JSON字符串
	logger.Info("Raw JSON data:", string(body))

	// 恢复请求体供后续使用
	ctx.Request.Body = io.NopCloser(strings.NewReader(string(body)))
}

// DebugContext 调试函数：查看Context中的关键信息
func DebugContext(ctx *gin.Context) {
	logger.Info("=== Context Debug Info ===")
	logger.Info("Method:", ctx.Request.Method)
	logger.Info("URL:", ctx.Request.URL.String())
	logger.Info("Content-Type:", ctx.GetHeader("Content-Type"))
	logger.Info("User-Agent:", ctx.GetHeader("User-Agent"))

	// 查看所有请求头
	logger.Info("All Headers:")
	for key, values := range ctx.Request.Header {
		for _, value := range values {
			logger.Info("  ", key, ":", value)
		}
	}

	// 查看URL参数
	if len(ctx.Request.URL.Query()) > 0 {
		logger.Info("Query Parameters:")
		for key, values := range ctx.Request.URL.Query() {
			for _, value := range values {
				logger.Info("  ", key, ":", value)
			}
		}
	}

	// 查看路径参数
	if len(ctx.Params) > 0 {
		logger.Info("Path Parameters:")
		for _, param := range ctx.Params {
			logger.Info("  ", param.Key, ":", param.Value)
		}
	}
}
