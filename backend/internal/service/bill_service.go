package service

import (
	"backend/api/response"
	"backend/internal/repository"
	"net/http"
	"time"

	"github.com/wonderivan/logger"
)

// CreateBill 创建账单
func CreateBill(userID int, billType int, amount float64, category, occurredAt, note, merchant, location, paymentMethod string) *Result[response.Bill] {
	// 解析时间
	parsedTime, err := time.Parse("2006-01-02 15:04:05", occurredAt)
	if err != nil {
		// 尝试解析仅日期格式
		parsedTime, err = time.Parse("2006-01-02", occurredAt)
		if err != nil {
			logger.Warn("Invalid time format:", err.Error())
			return ResultFailed[response.Bill](http.StatusBadRequest, "时间格式错误")
		}
	}

	// 根据分类名称获取分类ID
	categoryModel, err := repository.GetCategoryByName(category)
	if err != nil {
		logger.Warn("Category not found:", err.Error())
		return ResultFailed[response.Bill](http.StatusBadRequest, "分类不存在")
	}

	// 验证分类类型与账单类型是否匹配
	if categoryModel.Type != billType {
		return ResultFailed[response.Bill](http.StatusBadRequest, "分类类型与账单类型不匹配")
	}

	// 创建账单记录
	recordID, err := repository.CreateTransactionRecord(userID, categoryModel.CategoryID, amount, parsedTime, note, merchant, location, paymentMethod)
	if err != nil {
		logger.Warn("Create transaction record error:", err.Error())
		if err.Error() == "用户不存在或未关联家庭" {
			return ResultFailed[response.Bill](http.StatusBadRequest, err.Error())
		}
		return ResultFailed[response.Bill](http.StatusInternalServerError, "Internal server error")
	}

	// 获取用户信息用于响应
	user, err := repository.GetUserByID(userID)
	if err != nil {
		logger.Warn("Get user error:", err.Error())
		return ResultFailed[response.Bill](http.StatusInternalServerError, "Internal server error")
	}

	// 构造响应
	bill := response.Bill{
		ID:            recordID,
		Type:          billType,
		Amount:        amount,
		Category:      category,
		OccurredAt:    parsedTime.Unix(),
		Note:          note,
		Member:        user.UserName,
		Merchant:      merchant,
		Location:      location,
		PaymentMethod: paymentMethod,
	}

	return ResultOK(bill)
}

// QueryBills 查询账单
func QueryBills(userID int, billType *int, startDate, endDate, category, member string) *Result[response.BillList] {
	// 获取用户家庭ID
	user, err := repository.GetUserByID(userID)
	if err != nil {
		logger.Warn("Get user error:", err.Error())
		return ResultFailed[response.BillList](http.StatusInternalServerError, "Internal server error")
	}

	if user.FamilyID == nil {
		return ResultFailed[response.BillList](http.StatusBadRequest, "用户未加入家庭")
	}

	// 查询账单记录
	records, err := repository.GetTransactionRecordsByFamily(*user.FamilyID, billType, startDate, endDate, category, member)
	if err != nil {
		logger.Warn("Query transaction records error:", err.Error())
		return ResultFailed[response.BillList](http.StatusInternalServerError, "Internal server error")
	}

	// 转换为响应格式
	bills := make([]response.Bill, 0, len(records))
	for _, record := range records {
		// 获取分类信息
		var categoryName string
		var billType int
		err = repository.DB.Table("Category").
			Select("categoryname, type").
			Where("categoryid = ?", record.CategoryID).
			Row().Scan(&categoryName, &billType)
		if err != nil {
			logger.Warn("Get category error:", err.Error())
			continue
		}

		// 获取用户名
		var userName string
		if record.UserID != nil {
			err = repository.DB.Table("Users").
				Select("username").
				Where("userid = ?", *record.UserID).
				Row().Scan(&userName)
			if err != nil {
				logger.Warn("Get username error:", err.Error())
				userName = "未知用户"
			}
		} else {
			userName = "家庭整体"
		}

		bill := response.Bill{
			ID:            record.TransactionRecordID,
			Type:          billType,
			Amount:        record.Amount,
			Category:      categoryName,
			OccurredAt:    record.OccurredAt.Unix(),
			Note:          stringValueFromPtr(record.Note),
			Member:        userName,
			Merchant:      stringValueFromPtr(record.Merchant),
			Location:      stringValueFromPtr(record.Location),
			PaymentMethod: stringValueFromPtr(record.PaymentMethod),
		}
		bills = append(bills, bill)
	}

	return ResultOK(response.BillList{Bills: bills})
}

// CreateRecurringBill 创建定期账单（这里简化处理，实际应该有专门的定期账单表）
func CreateRecurringBill(userID int, billType int, amount float64, category, occurredAt, note, interval string) *Result[response.RecurringBill] {
	// 解析时间
	parsedTime, err := time.Parse("2006-01-02 15:04:05", occurredAt)
	if err != nil {
		// 尝试解析仅日期格式
		parsedTime, err = time.Parse("2006-01-02", occurredAt)
		if err != nil {
			logger.Warn("Invalid time format:", err.Error())
			return ResultFailed[response.RecurringBill](http.StatusBadRequest, "时间格式错误")
		}
	}

	// 验证周期参数
	if interval != "daily" && interval != "weekly" && interval != "monthly" {
		return ResultFailed[response.RecurringBill](http.StatusBadRequest, "周期参数错误，支持：daily, weekly, monthly")
	}

	// 根据分类名称获取分类ID
	categoryModel, err := repository.GetCategoryByName(category)
	if err != nil {
		logger.Warn("Category not found:", err.Error())
		return ResultFailed[response.RecurringBill](http.StatusBadRequest, "分类不存在")
	}

	// 验证分类类型与账单类型是否匹配
	if categoryModel.Type != billType {
		return ResultFailed[response.RecurringBill](http.StatusBadRequest, "分类类型与账单类型不匹配")
	}

	// 创建第一条账单记录
	recordID, err := repository.CreateTransactionRecord(userID, categoryModel.CategoryID, amount, parsedTime, note, "", "", "")
	if err != nil {
		logger.Warn("Create transaction record error:", err.Error())
		if err.Error() == "用户不存在或未关联家庭" {
			return ResultFailed[response.RecurringBill](http.StatusBadRequest, err.Error())
		}
		return ResultFailed[response.RecurringBill](http.StatusInternalServerError, "Internal server error")
	}

	// 构造响应（这里简化处理，实际应该存储定期任务信息）
	recurringBill := response.RecurringBill{
		ID:         recordID,
		Type:       billType,
		Amount:     amount,
		Category:   category,
		OccurredAt: parsedTime.Unix(),
		Note:       note,
		Interval:   interval,
	}

	return ResultOK(recurringBill)
}

// QueryRecurringBills 查询定期账单（简化处理）
func QueryRecurringBills(userID int) *Result[response.RecurringBillList] {
	// 这里简化处理，实际应该有专门的定期账单表
	// 返回空列表
	return ResultOK(response.RecurringBillList{Bills: []response.RecurringBill{}})
}

// GetIncomeStats 获取收入统计
func GetIncomeStats(userID int, startDate, endDate, category string) *Result[response.Stats] {
	// 获取统计数据
	stats, err := repository.GetFamilyFinanceStats(userID, startDate, endDate)
	if err != nil {
		logger.Warn("Get finance stats error:", err.Error())
		return ResultFailed[response.Stats](http.StatusInternalServerError, "Internal server error")
	}

	var totalAmount float64
	for _, stat := range stats {
		// 只统计收入（type=1）
		if incomeType, ok := stat["income_or_expense"].(int64); ok && incomeType == 1 {
			if amount, ok := stat["total_amount"].(float64); ok {
				totalAmount += amount
			}
		}
	}

	return ResultOK(response.Stats{
		Amount:   totalAmount,
		Category: category,
	})
}

// GetExpenseStats 获取支出统计
func GetExpenseStats(userID int, startDate, endDate, category string) *Result[response.Stats] {
	// 获取统计数据
	stats, err := repository.GetFamilyFinanceStats(userID, startDate, endDate)
	if err != nil {
		logger.Warn("Get finance stats error:", err.Error())
		return ResultFailed[response.Stats](http.StatusInternalServerError, "Internal server error")
	}

	var totalAmount float64
	for _, stat := range stats {
		// 只统计支出（type=0）
		if expenseType, ok := stat["income_or_expense"].(int64); ok && expenseType == 0 {
			if amount, ok := stat["total_amount"].(float64); ok {
				totalAmount += amount
			}
		}
	}

	return ResultOK(response.Stats{
		Amount:   totalAmount,
		Category: category,
	})
}

// QueryBudget 查询预算
func QueryBudget(userID int, startDate, category string) *Result[response.Budget] {
	// 获取用户家庭ID
	user, err := repository.GetUserByID(userID)
	if err != nil {
		logger.Warn("Get user error:", err.Error())
		return ResultFailed[response.Budget](http.StatusInternalServerError, "Internal server error")
	}

	if user.FamilyID == nil {
		return ResultFailed[response.Budget](http.StatusBadRequest, "用户未加入家庭")
	}

	// 查询预算
	budget, err := repository.GetBudgetByFamilyAndTime(*user.FamilyID, startDate)
	if err != nil {
		logger.Warn("Get budget error:", err.Error())
		return ResultFailed[response.Budget](http.StatusNotFound, "预算不存在")
	}

	return ResultOK(response.Budget{
		StartDate: budget.Time,
		Amount:    budget.Amount,
		Category:  category,
		Note:      "",
	})
}

// stringValueFromPtr 辅助函数：从字符串指针获取值，nil返回空字符串
func stringValueFromPtr(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

// DeleteBill 删除账单
func DeleteBill(billID int) *Result[string] {
	// 调用repository层删除账单
	err := repository.DeleteTransactionRecord(billID)
	if err != nil {
		logger.Warn("Delete bill error:", err.Error())
		if err.Error() == "账单不存在" {
			return ResultFailed[string](http.StatusNotFound, "账单不存在")
		}
		return ResultFailed[string](http.StatusInternalServerError, "Internal server error")
	}

	return ResultOK("账单删除成功")
}
