package middleware

import (
	"backend/pkg/config"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

type AuthUser struct {
	ID int
	jwt.StandardClaims
}

// JWTAuth JWT认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userToken := ctx.GetHeader("Authorization")
		if userToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "请提供认证token",
			})
			ctx.Abort()
			return
		}

		// 移除 "Bearer " 前缀
		if len(userToken) > 7 && userToken[:7] == "Bearer " {
			userToken = userToken[7:]
		}

		claims := ParseTokenWithVerify(userToken)
		if claims == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "无效的token",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.ID)
		ctx.Next()
	}
}

func ParseTokenWithVerify(tokenString string) *AuthUser {
	token, err := jwt.ParseWithClaims(tokenString, &AuthUser{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.App.JwtSecret), nil
	})

	if err != nil {
		logger.Warn("(ParseTokenWithVerify)Error when parse token: %v", err.Error())
		return nil
	}

	if !token.Valid {
		logger.Warn("(ParseTokenWithVerify)Invalid token")
		return nil
	}

	claims, ok := token.Claims.(*AuthUser)
	if ok {
		logger.Trace("(ParseTokenWithVerify)Parse token successfully.")
		return claims
	} else {
		logger.Warn("(ParseTokenWithVerify)Invalid token claims")
		return nil
	}
}
