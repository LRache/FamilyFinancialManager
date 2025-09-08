package handler

import (
	"backend/api/request"
	"backend/api/response"
	"backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateFamily(ctx *gin.Context) {
	var req request.CreateFamily
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	userIDInterface, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "未登录")
		return
	}

	userID, ok := userIDInterface.(int)
	if !ok {
		response.BadRequest(ctx, "无效的用户ID")
		return
	}

	result := service.CreateFamily(userID, req.FamilyName)

	ResultToResponse(ctx, result, result.Data)
}

func InviteUserToFamily(ctx *gin.Context) {
	var req request.InviteUser
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	userIDInterface, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "未登录")
		return
	}

	userID, ok := userIDInterface.(int)
	if !ok {
		response.BadRequest(ctx, "无效的用户ID")
		return
	}

	result := service.InviteUserToFamily(userID, req.Username)

	ResultToResponse(ctx, result, result.Data)
}

func GetFamilyMembers(ctx *gin.Context) {
	userIDInterface, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "未登录")
		return
	}

	userID, ok := userIDInterface.(int)
	if !ok {
		response.BadRequest(ctx, "无效的用户ID")
		return
	}

	result := service.GetFamilyMembers(userID)

	ResultToResponse(ctx, result, result.Data)
}

func GetFamilyInfo(ctx *gin.Context) {
	userIDInterface, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "未登录")
		return
	}

	userID, ok := userIDInterface.(int)
	if !ok {
		response.BadRequest(ctx, "无效的用户ID")
		return
	}

	result := service.GetFamilyInfo(userID)

	ResultToResponse(ctx, result, result.Data)
}

func GetFamilyInfoByID(ctx *gin.Context) {
	familyIDStr := ctx.Param("familyid")
	_, err := strconv.Atoi(familyIDStr)
	if err != nil {
		response.BadRequest(ctx, "无效的家庭ID")
		return
	}

	userIDInterface, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "未登录")
		return
	}

	userID, ok := userIDInterface.(int)
	if !ok {
		response.BadRequest(ctx, "无效的用户ID")
		return
	}

	result := service.GetFamilyInfo(userID)

	ResultToResponse(ctx, result, result.Data)
}
