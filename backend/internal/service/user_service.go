package service

import (
	"backend/api/response"
	"backend/internal/repository"
	"net/http"

	"github.com/wonderivan/logger"
)

func RegisterUser(username string, password string, email string) *Result[response.UserLogin] {
	ok, err := repository.UserExists(username)
	if err != nil {
		logger.Warn("Internal server error", err.Error())
		return ResultFailed[response.UserLogin](http.StatusInternalServerError, "Internal server error")
	}

	if ok {
		return ResultFailed[response.UserLogin](http.StatusConflict, "User already exists")
	}

	id, err := repository.CreateUser(username, password, email)
	if err != nil {
		logger.Warn("Internal server error", err.Error())
		return ResultFailed[response.UserLogin](http.StatusInternalServerError, "Internal server error")
	}

	token := GenerateAuthToken(id)

	return ResultOK(response.UserLogin{
		Token: token,
	})
}

func UserLogin(username string, password string) *Result[response.UserLogin] {
	id, err := repository.UserLogin(username, password)
	if err != nil {
		logger.Warn("Internal server error", err.Error())
		return ResultFailed[response.UserLogin](http.StatusInternalServerError, "Internal server error")
	}

	if id == -1 {
		return ResultFailed[response.UserLogin](http.StatusUnauthorized, "Invalid username or password")
	}

	token := GenerateAuthToken(id)

	return ResultOK(response.UserLogin{
		Token: token,
	})
}
