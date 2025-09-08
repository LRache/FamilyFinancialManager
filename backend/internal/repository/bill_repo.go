package repository

import (
	"backend/internal/model"
	"errors"
	"strings"
	"time"
)

// CreateTransactionRecord 调用存储过程添加收支记录
func CreateTransactionRecord(userID int, categoryID int, amount float64, occurredAt time.Time, note, merchant, location, paymentMethod string) (int, error) {
	// 调用存储过程 AddTransactionRecord
	result := DB.Exec("CALL AddTransactionRecord(?, ?, ?, ?, ?, ?, ?, ?)",
		userID, categoryID, amount, occurredAt,
		stringPtrFromString(note),
		stringPtrFromString(merchant),
		stringPtrFromString(location),
		stringPtrFromString(paymentMethod))

	if result.Error != nil {
		// 检查是否是MySQL 45000异常（业务逻辑错误）
		errorMsg := result.Error.Error()
		if strings.Contains(errorMsg, "Error 1644") {
			if strings.Contains(errorMsg, "分类不存在") {
				return 0, errors.New("分类不存在")
			}
			if strings.Contains(errorMsg, "用户不存在或未关联家庭") {
				return 0, errors.New("用户不存在或未关联家庭")
			}
		}
		return 0, result.Error
	}

	// 获取最后插入的记录ID
	var lastID int
	err := DB.Raw("SELECT LAST_INSERT_ID()").Scan(&lastID).Error
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

// GetTransactionRecordsByFamily 根据家庭ID查询收支记录
func GetTransactionRecordsByFamily(familyID int, billType *int, startDate, endDate, category, member string) ([]model.TransactionRecord, error) {
	var records []model.TransactionRecord

	query := DB.Table("TransactionRecord t").
		Select("t.*").
		Joins("JOIN Category c ON t.categoryid = c.categoryid").
		Joins("LEFT JOIN Users u ON t.userid = u.userid").
		Where("t.familyid = ?", familyID)

	// 添加类型过滤
	if billType != nil {
		query = query.Where("c.type = ?", *billType)
	}

	// 添加日期范围过滤
	if startDate != "" {
		query = query.Where("DATE(t.occurred_at) >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("DATE(t.occurred_at) <= ?", endDate)
	}

	// 添加分类过滤
	if category != "" {
		query = query.Where("c.categoryname = ?", category)
	}

	// 添加成员过滤
	if member != "" {
		query = query.Where("u.username = ?", member)
	}

	err := query.Find(&records).Error
	return records, err
}

// EditTransactionRecord 调用存储过程修改收支记录
func EditTransactionRecord(recordID, userID, categoryID int, amount float64, occurredAt time.Time, note, merchant, location, paymentMethod string) error {
	// 调用存储过程 EditTransactionRecord
	result := DB.Exec("CALL EditTransactionRecord(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		recordID, userID, categoryID, amount, occurredAt,
		stringPtrFromString(note),
		stringPtrFromString(merchant),
		stringPtrFromString(location),
		stringPtrFromString(paymentMethod))

	if result.Error != nil {
		// 检查是否是MySQL 45000异常（业务逻辑错误）
		errorMsg := result.Error.Error()
		if strings.Contains(errorMsg, "Error 1644") {
			if strings.Contains(errorMsg, "账单不存在") {
				return errors.New("账单不存在")
			}
			if strings.Contains(errorMsg, "用户不存在或不属于同一家庭") {
				return errors.New("用户不存在或不属于同一家庭")
			}
			if strings.Contains(errorMsg, "分类不存在") {
				return errors.New("分类不存在")
			}
		}
		return result.Error
	}

	return nil
}

// DeleteTransactionRecord 调用存储过程删除收支记录
func DeleteTransactionRecord(recordID int) error {
	// 调用存储过程 DeleteTransactionRecord
	result := DB.Exec("CALL DeleteTransactionRecord(?)", recordID)

	if result.Error != nil {
		// 检查是否是MySQL 45000异常（业务逻辑错误）
		errorMsg := result.Error.Error()
		if strings.Contains(errorMsg, "Error 1644") {
			if strings.Contains(errorMsg, "账单不存在") {
				return errors.New("账单不存在")
			}
		}
		return result.Error
	}

	return nil
}

// GetCategoryByName 根据分类名称获取分类信息
func GetCategoryByName(categoryName string) (*model.Category, error) {
	var category model.Category
	err := DB.Where("categoryname = ?", categoryName).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetAllCategories 获取所有分类
func GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	err := DB.Find(&categories).Error
	return categories, err
}

// GetFamilyFinanceStats 调用存储过程获取家庭收支统计
func GetFamilyFinanceStats(userID int, startDate, endDate string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	// 调用存储过程 GetFamilyFinanceByUser
	rows, err := DB.Raw("CALL GetFamilyFinanceByUser(?, ?, ?)", userID, startDate, endDate).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// 读取结果
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}

		result := make(map[string]interface{})
		for i, col := range columns {
			result[col] = values[i]
		}
		results = append(results, result)
	}

	return results, nil
}

// GetBudgetByFamilyAndTime 获取家庭预算
func GetBudgetByFamilyAndTime(familyID int, time string) (*model.Budget, error) {
	var budget model.Budget
	err := DB.Where("familyid = ? AND time = ?", familyID, time).First(&budget).Error
	if err != nil {
		return nil, err
	}
	return &budget, nil
}

// stringPtrFromString 辅助函数：将字符串转换为字符串指针，空字符串返回nil
func stringPtrFromString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
