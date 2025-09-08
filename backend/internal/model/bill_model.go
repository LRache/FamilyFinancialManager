package model

import "time"

// TransactionRecord 收支记录模型
type TransactionRecord struct {
	TransactionRecordID int       `gorm:"column:transactionrecordid;primaryKey;autoIncrement" json:"id"`
	FamilyID            int       `gorm:"column:familyid;not null" json:"familyid"`
	UserID              *int      `gorm:"column:userid" json:"userid"`
	CategoryID          int       `gorm:"column:categoryid;not null" json:"categoryid"`
	Amount              float64   `gorm:"column:amount;not null" json:"amount"`
	OccurredAt          time.Time `gorm:"column:occurred_at;not null" json:"occurred_at"`
	Note                *string   `gorm:"column:note" json:"note"`
	Merchant            *string   `gorm:"column:merchant" json:"merchant"`
	Location            *string   `gorm:"column:location" json:"location"`
	PaymentMethod       *string   `gorm:"column:paymentmethod" json:"payment_method"`
}

func (TransactionRecord) TableName() string {
	return "TransactionRecord"
}

// Category 分类模型
type Category struct {
	CategoryID   int     `gorm:"column:categoryid;primaryKey;autoIncrement" json:"categoryid"`
	CategoryName string  `gorm:"column:categoryname;not null" json:"categoryname"`
	Type         int     `gorm:"column:type;not null" json:"type"` // 1=收入, 0=支出
	Note         *string `gorm:"column:note" json:"note"`
}

func (Category) TableName() string {
	return "Category"
}

// Budget 预算模型
type Budget struct {
	FamilyID int     `gorm:"column:familyid;primaryKey" json:"familyid"`
	Time     string  `gorm:"column:time;primaryKey" json:"time"` // 预算对应的时间（按月或按年）
	Amount   float64 `gorm:"column:amount;not null" json:"amount"`
}

func (Budget) TableName() string {
	return "Budget"
}

// RecurringBill 定期账单模型（用于API中的定期收支）
type RecurringBill struct {
	ID         int       `json:"id"`
	Type       int       `json:"type"`        // 账单类型（0表示支出，1表示收入）
	Amount     float64   `json:"amount"`      // 账单数额
	Category   string    `json:"category"`    // 账单类别
	OccurredAt time.Time `json:"occurred_at"` // 发生时间
	Note       string    `json:"note"`        // 备注
	Interval   string    `json:"interval"`    // 账单周期（支持 daily, weekly, monthly）
}
