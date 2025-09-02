package service

import (
	"backend/api/response"
	"backend/internal/repository"
	"net/http"

	"github.com/wonderivan/logger"
)

func RegisterUser(username string, password string) *Result {
	ok, err := repository.UserExists(username)
	if err != nil {
		logger.Warn("Internal server error", err.Error())
		return &Result{Code: http.StatusInternalServerError, Message: "Internal server error"}
	}

	if ok {
		return &Result{Code: http.StatusConflict, Message: "User already exists"}
	}

	id, err := repository.CreateUser(username, password)
	if err != nil {
		logger.Warn("Internal server error", err.Error())
		return &Result{Code: http.StatusInternalServerError, Message: "Internal server error"}
	}

	return &Result{Code: http.StatusOK, Message: "User registered successfully", Data: &response.UserLogin{
		Username: username,
		ID:       id,
		Token:    GenerateAuthToken(id),
	}}
}

func UserLogin(username string, password string) *Result {
	id, err := repository.UserLogin(username, password)
	if err != nil {
		logger.Warn("Internal server error", err.Error())
		return &Result{Code: http.StatusInternalServerError, Message: "Internal server error"}
	}

	if id == -1 {
		return &Result{Code: http.StatusUnauthorized, Message: "Invalid username or password"}
	}

	return &Result{Code: http.StatusOK, Message: "User logged in successfully"}
}
