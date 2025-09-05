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
		familyGroup.POST("", handler.CreateFamily)               // 创建家庭 POST /api/v1/family
		familyGroup.POST("/members", handler.InviteUserToFamily) // 邀请用户 POST /api/v1/family/members
		familyGroup.GET("/members", handler.GetFamilyMembers)    // 获取成员列表 GET /api/v1/family/members
	}

	// 账单管理（需要认证）
	// TODO: 实现账单管理handlers
	/*
		billsGroup := root.Group("/bills")
		billsGroup.Use(middleware.JWTAuth())
		{
			billsGroup.POST("", handler.CreateBill)              // 上传账单 POST /api/v1/bills
			billsGroup.GET("", handler.QueryBills)               // 查询账单 GET /api/v1/bills
			billsGroup.POST("/recurring", handler.AddRecurringBill) // 添加定期收支 POST /api/v1/bills/recurring
			billsGroup.GET("/recurring", handler.QueryRecurringBills) // 查询定期收支 GET /api/v1/bills/recurring
			billsGroup.GET("/income/stats", handler.GetIncomeStats)   // 查询收入统计 GET /api/v1/bills/income/stats
		}
	*/

	// 预算管理（需要认证）
	// TODO: 实现预算管理handlers
	/*
		budgetGroup := root.Group("/budget")
		budgetGroup.Use(middleware.JWTAuth())
		{
			budgetGroup.POST("", handler.SetBudget)      // 设置预算 POST /api/v1/budget
			budgetGroup.GET("", handler.QueryBudget)     // 查询预算 GET /api/v1/budget
		}
	*/

	// 支出统计（需要认证）
	// TODO: 实现统计handlers
	/*
		statsGroup := root.Group("/outcomes")
		statsGroup.Use(middleware.JWTAuth())
		{
			statsGroup.GET("/stats", handler.GetOutcomeStats) // 查询支出统计 GET /api/v1/outcomes/stats
		}
	*/
}
