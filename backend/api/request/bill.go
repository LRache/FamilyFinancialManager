package request

// CreateBill 上传账单请求
type CreateBill struct {
	Type          int     `form:"type" json:"type" binding:"required"`               // 账单类型（0表示支出，1表示收入）
	Amount        float64 `form:"amount" json:"amount" binding:"required"`           // 账单数额
	Category      string  `form:"category" json:"category" binding:"required"`       // 账单类别
	OccurredAt    string  `form:"occurred_at" json:"occurred_at" binding:"required"` // 发生时间
	Note          string  `form:"note" json:"note"`                                  // 备注
	Merchant      string  `form:"merchant" json:"merchant"`                          // 商家名称
	Location      string  `form:"location" json:"location"`                          // 消费地点
	PaymentMethod string  `form:"payment_method" json:"payment_method"`              // 支付方式
}

// QueryBill 查询账单请求
type QueryBill struct {
	Type      *int   `form:"type" json:"type"`             // 账单类型（0表示支出，1表示收入）
	StartDate string `form:"start_date" json:"start_date"` // 开始日期
	EndDate   string `form:"end_date" json:"end_date"`     // 结束日期
	Category  string `form:"category" json:"category"`     // 账单类别
	Member    string `form:"member" json:"member"`         // 家庭成员
}

// CreateRecurringBill 添加定期收支请求
type CreateRecurringBill struct {
	Type       int     `form:"type" json:"type" binding:"required"`               // 账单类型（0表示支出，1表示收入）
	Amount     float64 `form:"amount" json:"amount" binding:"required"`           // 账单数额
	Category   string  `form:"category" json:"category" binding:"required"`       // 账单类别
	OccurredAt string  `form:"occurred_at" json:"occurred_at" binding:"required"` // 发生时间
	Note       string  `form:"note" json:"note"`                                  // 备注
	Interval   string  `form:"interval" json:"interval" binding:"required"`       // 账单周期（支持 daily, weekly, monthly）
}

// QueryBudget 查询预算请求
type QueryBudget struct {
	StartDate string `form:"start_date" json:"start_date"` // 开始日期
	Category  string `form:"category" json:"category"`     // 预算类别
}

// QueryStats 查询统计请求
type QueryStats struct {
	StartDate string `form:"start_date" json:"start_date"` // 开始日期
	EndDate   string `form:"end_date" json:"end_date"`     // 结束日期
	Category  string `form:"category" json:"category"`     // 类别
}
