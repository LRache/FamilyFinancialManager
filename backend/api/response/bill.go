package response

// Bill 账单响应对象
type Bill struct {
	ID            int     `json:"id"`             // 账单ID
	Type          int     `json:"type"`           // 账单类型
	Amount        float64 `json:"amount"`         // 账单数额
	Category      string  `json:"category"`       // 账单类别
	OccurredAt    int64   `json:"occurred_at"`    // 发生时间戳
	Note          string  `json:"note"`           // 备注
	Member        string  `json:"member"`         // 家庭成员
	Merchant      string  `json:"merchant"`       // 商家名称
	Location      string  `json:"location"`       // 消费地点
	PaymentMethod string  `json:"payment_method"` // 支付方式
}

// BillList 账单列表响应
type BillList struct {
	Bills []Bill `json:"bills"`
}

// RecurringBill 定期账单响应对象
type RecurringBill struct {
	ID         int     `json:"id"`          // 账单ID
	Type       int     `json:"type"`        // 账单类型
	Amount     float64 `json:"amount"`      // 账单数额
	Category   string  `json:"category"`    // 账单类别
	OccurredAt int64   `json:"occurred_at"` // 发生时间戳
	Note       string  `json:"note"`        // 备注
	Interval   string  `json:"interval"`    // 账单周期
}

// RecurringBillList 定期账单列表响应
type RecurringBillList struct {
	Bills []RecurringBill `json:"bills"`
}

// Budget 预算响应对象
type Budget struct {
	StartDate string  `json:"start_date"` // 开始日期
	Amount    float64 `json:"amount"`     // 预算金额
	Category  string  `json:"category"`   // 预算类别
	Note      string  `json:"note"`       // 备注
}

// Stats 统计响应对象
type Stats struct {
	Amount   float64 `json:"amount"`   // 金额
	Category string  `json:"category"` // 类别
}
