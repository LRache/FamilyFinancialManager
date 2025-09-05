package handler

import (
	"backend/api/request"
	"backend/api/response"
	"backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateFamily 创建家庭
func CreateFamily(ctx *gin.Context) {
	var req request.CreateFamily
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	// 从JWT token中获取用户ID（假设middleware已经设置了user_id）
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

// InviteUserToFamily 邀请用户加入家庭
func InviteUserToFamily(ctx *gin.Context) {
	var req request.InviteUser
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	// 从JWT token中获取用户ID
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

// GetFamilyMembers 获取家庭成员列表
func GetFamilyMembers(ctx *gin.Context) {
	// 从JWT token中获取用户ID
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

// SetFamilyBudget 设置家庭预算
func SetFamilyBudget(ctx *gin.Context) {
	var req request.SetBudget
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	// 从JWT token中获取用户ID
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

	result := service.SetFamilyBudget(userID, req.StartDate, req.Amount)

	ResultToResponse(ctx, result, gin.H{"message": result.Data})
}

// GetFamilyInfo 获取家庭信息
func GetFamilyInfo(ctx *gin.Context) {
	// 从JWT token中获取用户ID
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

// GetFamilyInfoByID 根据家庭ID获取家庭信息（管理员功能）
func GetFamilyInfoByID(ctx *gin.Context) {
	familyIDStr := ctx.Param("familyid")
	_, err := strconv.Atoi(familyIDStr)
	if err != nil {
		response.BadRequest(ctx, "无效的家庭ID")
		return
	}

	// 从JWT token中获取用户ID
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

	// 这里可以添加权限检查，确保用户有权限查看该家庭信息
	result := service.GetFamilyInfo(userID)

	ResultToResponse(ctx, result, result.Data)
}
