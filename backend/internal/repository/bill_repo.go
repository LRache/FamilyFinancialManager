package repository

import (
	"backend/internal/model"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/wonderivan/logger"
)

// CreateTransactionRecord 调用存储过程添加收支记录，返回记录ID和是否超支标志
func CreateTransactionRecord(userID int, categoryID int, amount float64, occurredAt time.Time, note, merchant, location, paymentMethod string) (int, bool, error) {
	var recordID int
	var exceedFlag sql.NullInt32

	err := DB.Exec("CALL AddTransactionRecord(?, ?, ?, ?, ?, ?, ?, ?, @exceed_flag)",
		userID, categoryID, amount, occurredAt,
		stringPtrFromString(note),
		stringPtrFromString(merchant),
		stringPtrFromString(location),
		stringPtrFromString(paymentMethod)).Error

	if err != nil {
		errorMsg := err.Error()
		if strings.Contains(errorMsg, "Error 1644") {
			if strings.Contains(errorMsg, "分类不存在") {
				return 0, false, errors.New("分类不存在")
			}
			if strings.Contains(errorMsg, "用户不存在或未关联家庭") {
				return 0, false, errors.New("用户不存在或未关联家庭")
			}
		}
		return 0, false, err
	}

	err = DB.Raw("SELECT @exceed_flag").Scan(&exceedFlag).Error
	if err != nil {
		return 0, false, err
	}

	logger.Info("Exceed flag value:", exceedFlag)

	err = DB.Raw("SELECT LAST_INSERT_ID()").Scan(&recordID).Error
	if err != nil {
		return 0, false, err
	}

	isOverBudget := exceedFlag.Valid && exceedFlag.Int32 == 1

	return recordID, isOverBudget, nil
}

// GetTransactionRecordsByFamily 根据家庭ID查询收支记录
func GetTransactionRecordsByFamily(familyID int, billType *int, startDate, endDate, category, member string) ([]model.TransactionRecord, error) {
	var records []model.TransactionRecord

	query := DB.Table("TransactionRecord t").
		Select("t.*").
		Joins("JOIN Category c ON t.categoryid = c.categoryid").
		Joins("LEFT JOIN Users u ON t.userid = u.userid").
		Where("t.familyid = ?", familyID)

	if billType != nil {
		query = query.Where("c.type = ?", *billType)
	}

	if startDate != "" {
		query = query.Where("DATE(t.occurred_at) >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("DATE(t.occurred_at) <= ?", endDate)
	}

	if category != "" {
		query = query.Where("c.categoryname = ?", category)
	}

	if member != "" {
		query = query.Where("u.username = ?", member)
	}

	err := query.Find(&records).Error
	return records, err
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

func GetFamilyFinanceStats(userID int, startDate, endDate string) ([]map[string]interface{}, error) {
	var results []map[string]any

	rows, err := DB.Raw("CALL GetFamilyFinanceByUser(?, ?, ?)", userID, startDate, endDate).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		values := make([]any, len(columns))
		valuePtrs := make([]any, len(columns))

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}

		result := make(map[string]any)
		for i, col := range columns {
			result[col] = values[i]
		}
		results = append(results, result)
	}

	return results, nil
}

func GetBudgetByFamilyAndTime(familyID int, time string) (*model.Budget, error) {
	var budget model.Budget
	err := DB.Where("familyid = ? AND time = ?", familyID, time).First(&budget).Error
	if err != nil {
		return nil, err
	}
	return &budget, nil
}

func stringPtrFromString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
