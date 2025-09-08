package handler

import (
	"backend/api/request"
	"backend/api/response"
	"backend/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

// CreateBill 上传账单
func CreateBill(ctx *gin.Context) {
	var req request.CreateBill
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

	result := service.CreateBill(userID, req.Type, req.Amount, req.Category, req.OccurredAt, req.Note, req.Merchant, req.Location, req.PaymentMethod)

	ResultToResponse(ctx, result, result.Data)
}

// QueryBills 查询账单
func QueryBills(ctx *gin.Context) {
	var req request.QueryBill
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

	result := service.QueryBills(userID, req.Type, req.StartDate, req.EndDate, req.Category, req.Member)

	ResultToResponse(ctx, result, result.Data)
}

// AddRecurringBill 添加定期收支
func AddRecurringBill(ctx *gin.Context) {
	var req request.CreateRecurringBill
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

	result := service.CreateRecurringBill(userID, req.Type, req.Amount, req.Category, req.OccurredAt, req.Note, req.Interval)

	ResultToResponse(ctx, result, result.Data)
}

// QueryRecurringBills 查询定期收支
func QueryRecurringBills(ctx *gin.Context) {
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

	result := service.QueryRecurringBills(userID)

	ResultToResponse(ctx, result, result.Data)
}

// GetIncomeStats 查询收入统计
func GetIncomeStats(ctx *gin.Context) {
	var req request.QueryStats
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

	result := service.GetIncomeStats(userID, req.StartDate, req.EndDate, req.Category)

	ResultToResponse(ctx, result, result.Data)
}

// GetExpenseStats 查询支出统计
func GetExpenseStats(ctx *gin.Context) {
	var req request.QueryStats
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

	result := service.GetExpenseStats(userID, req.StartDate, req.EndDate, req.Category)

	ResultToResponse(ctx, result, result.Data)
}

// SetBudget 设置预算（复用已有的家庭预算设置功能）
func SetBudget(ctx *gin.Context) {
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

// DeleteBill 删除账单
func DeleteBill(ctx *gin.Context) {
	// 从URL参数中获取账单ID
	billIDStr := ctx.Param("id")
	if billIDStr == "" {
		response.BadRequest(ctx, "账单ID不能为空")
		return
	}

	// 转换账单ID为整数
	var billID int
	if _, err := fmt.Sscanf(billIDStr, "%d", &billID); err != nil {
		response.BadRequest(ctx, "无效的账单ID")
		return
	}

	// 从JWT token中获取用户ID（可以用于后续权限验证）
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
