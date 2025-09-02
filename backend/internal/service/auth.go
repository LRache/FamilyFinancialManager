package service

import (
	"backend/internal/middleware"
	"backend/pkg/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/wonderivan/logger"
)

func GenerateAuthToken(userId int) string {
	claims := &middleware.AuthUser{
		ID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(config.App.JwtSecret))
	if err != nil {
		logger.Warn("Failed to generate auth token", err.Error())
		return ""
	}
	return s
}
