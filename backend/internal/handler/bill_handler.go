package handler

import (
	"backend/api/request"
	"backend/api/response"
	"backend/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

func CreateBill(ctx *gin.Context) {
	// DebugContext(ctx)
	// DebugRequestBody(ctx)

	var req request.CreateBill
	err := ctx.ShouldBind(&req)
	if err != nil {
		logger.Warn("Failed to bind request:", err.Error())
		logger.Warn("Request binding failed. Content-Type:", ctx.GetHeader("Content-Type"))
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
		logger.Warn("Invalid user ID type")
		response.BadRequest(ctx, "无效的用户ID")
		return
	}

	result := service.CreateBill(userID, *req.Type, req.Amount, req.Category, req.OccurredAt, req.Note, req.Merchant, req.Location, req.PaymentMethod)

	ResultToResponse(ctx, result, result.Data)
}

func QueryBills(ctx *gin.Context) {
	var req request.QueryBill
	err := ctx.ShouldBindQuery(&req)
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

	result := service.QueryBills(userID, req.Type, req.StartDate, req.EndDate, req.Category, req.Member)

	ResultToResponse(ctx, result, result.Data)
}

func GetIncomeStats(ctx *gin.Context) {
	var req request.QueryStats
	err := ctx.ShouldBindQuery(&req)
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

	result := service.GetIncomeStats(userID, req.StartDate, req.EndDate, req.Category)

	ResultToResponse(ctx, result, result.Data)
}

func GetExpenseStats(ctx *gin.Context) {
	var req request.QueryStats
	err := ctx.ShouldBindQuery(&req)
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

	result := service.GetExpenseStats(userID, req.StartDate, req.EndDate, req.Category)

	ResultToResponse(ctx, result, result.Data)
}

func SetBudget(ctx *gin.Context) {
	var req request.SetBudget
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

	result := service.SetFamilyBudget(userID, req.StartDate, req.Amount)

	ResultToResponse(ctx, result, gin.H{"message": result.Data})
}

// QueryBudget 查询预算
func QueryBudget(ctx *gin.Context) {
	var req request.QueryBudget
	err := ctx.ShouldBindQuery(&req)
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

	result := service.QueryBudget(userID, req.StartDate, req.Category)

	ResultToResponse(ctx, result, result.Data)
}

func DeleteBill(ctx *gin.Context) {
	billIDStr := ctx.Param("id")
	if billIDStr == "" {
		response.BadRequest(ctx, "账单ID不能为空")
		return
	}

	var billID int
	if _, err := fmt.Sscanf(billIDStr, "%d", &billID); err != nil {
		response.BadRequest(ctx, "无效的账单ID")
		return
	}

	userIDInterface, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "未登录")
		return
	}

	_, ok := userIDInterface.(int)
	if !ok {
		response.BadRequest(ctx, "无效的用户ID")
		return
	}

	// 调用service层删除账单
	result := service.DeleteBill(billID)

	ResultToResponse(ctx, result, gin.H{"message": result.Data})
}
