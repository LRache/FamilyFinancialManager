package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
	// 允许的请求源（根据需求修改，* 表示允许所有）
	origin := c.Request.Header.Get("Origin")
	if origin != "" {
		c.Header("Access-Control-Allow-Origin", origin) // 允许当前源
		// 若需允许所有源，可改为：c.Header("Access-Control-Allow-Origin", "*")
	}

	// 允许的请求方法
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	// 允许的请求头（根据实际需求添加，如 Authorization、Token 等）
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// 是否允许携带 Cookie（若为 true，Access-Control-Allow-Origin 不能设为 *）
	c.Header("Access-Control-Allow-Credentials", "true")

	// 处理 OPTIONS 预检请求（浏览器在跨域请求前可能发送 OPTIONS 检查服务器是否允许）
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent) // 204 表示成功但无内容
		return
	}

	// 继续处理后续请求
	c.Next()
}
