package router

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/pkg/config"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	root := router.Group(config.App.RouterPrefix)

	authGroup := root.Group("/auth")
	authGroup.POST("/register", handler.RegisterUser)
	authGroup.POST("/login", handler.UserLogin)

	familyGroup := root.Group("/family")
	familyGroup.Use(middleware.JWTAuth())
	{
		familyGroup.POST("", handler.CreateFamily)
		familyGroup.POST("/members", handler.InviteUserToFamily)
		familyGroup.GET("/members", handler.GetFamilyMembers)
	}

	// 账单管理（需要认证）
	billsGroup := root.Group("/bills")
	billsGroup.Use(middleware.JWTAuth())
	{
		billsGroup.POST("", handler.CreateBill)
		billsGroup.GET("", handler.QueryBills)
		billsGroup.DELETE("/:id", handler.DeleteBill)
		billsGroup.GET("/income", handler.GetIncomeStats)
	}

	// 预算管理（需要认证）
	budgetGroup := root.Group("/budget")
	budgetGroup.Use(middleware.JWTAuth())
	{
		budgetGroup.POST("", handler.SetBudget)  // 设置预算 POST /api/v1/budget
		budgetGroup.GET("", handler.QueryBudget) // 查询预算 GET /api/v1/budget
	}

	// 支出统计（需要认证）
	outcomeGroup := root.Group("/outcome")
	outcomeGroup.Use(middleware.JWTAuth())
	{
		outcomeGroup.GET("", handler.GetExpenseStats) // 查询支出统计 GET /api/v1/outcome
	}
}
