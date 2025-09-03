package handler

import (
	"backend/api/request"
	"backend/api/response"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var req request.UserLogin
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	result := service.RegisterUser(req.Username, req.Password)

	ResultToResponse(ctx, result, result.Data)
}

func UserLogin(ctx *gin.Context) {
	var req request.UserLogin
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	result := service.UserLogin(req.Username, req.Password)

	ResultToResponse(ctx, result, result.Data)
}
