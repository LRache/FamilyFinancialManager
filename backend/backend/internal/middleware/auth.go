package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

type AuthUser struct {
	ID int
	jwt.StandardClaims
}

func AuthVerify(ctx *gin.Context) {
	userToken := ctx.GetHeader("auth")
	if userToken == "" {
		ctx.Set("auth", nil)
		return
	}

	ctx.Set("auth", ParseToken(userToken))
}

func ParseToken(tokenString string) *AuthUser {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &AuthUser{})
	if err != nil {
		logger.Warn("(ParseToken)Error when parse token, invalid token: err = \"%v\", tokenString = \"%v\"", err.Error())
		return nil
	}
	claims, ok := token.Claims.(*AuthUser)
	if ok {
		logger.Trace("(ParseToken)Parse token successfully.")
		return claims
	} else {
		logger.Warn("(ParseToken)Invalid token, tokenString = \"%v\"", tokenString)
		return nil
	}
}
