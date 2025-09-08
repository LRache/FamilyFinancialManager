package repository

import (
	"backend/internal/model"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// CreateFamily 调用存储过程创建家庭
func CreateFamily(userID int, familyName string) (int, error) {
	// 调用存储过程 sp_create_family
	result := DB.Exec("CALL sp_create_family(?, ?, @familyid)", userID, familyName)
	if result.Error != nil {
		// 检查是否是MySQL 45000异常（业务逻辑错误）
		if strings.Contains(result.Error.Error(), "Error 1644") ||
			strings.Contains(result.Error.Error(), "该用户已属于某个家庭，无法再次创建家庭") {
			return 0, errors.New("该用户已属于某个家庭，无法再次创建家庭")
		}
		return 0, result.Error
	}

	// 获取输出参数 familyid
	var familyID int
	err := DB.Raw("SELECT @familyid").Scan(&familyID).Error
	if err != nil {
		return 0, err
	}

	return familyID, nil
}

// GetFamilyByID 根据家庭ID获取家庭信息
func GetFamilyByID(familyID int) (*model.Family, error) {
	var family model.Family
	err := DB.Where("familyid = ?", familyID).First(&family).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &family, nil
}

// GetFamilyMembers 获取家庭成员列表
func GetFamilyMembers(familyID int) ([]model.User, error) {
	var users []model.User
	err := DB.Where("familyid = ?", familyID).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// InviteUserToFamily 调用存储过程邀请用户加入家庭
func InviteUserToFamily(inviterID int, inviteeID int) error {
	// 调用存储过程 sp_invite_user_to_family
	result := DB.Exec("CALL sp_invite_user_to_family(?, ?)", inviterID, inviteeID)
	if result.Error != nil {
		// 检查是否是MySQL 45000异常（业务逻辑错误）
		errorMsg := result.Error.Error()
		if strings.Contains(errorMsg, "Error 1644") {
			if strings.Contains(errorMsg, "只有家庭管理员才能邀请用户加入家庭") {
				return errors.New("只有家庭管理员才能邀请用户加入家庭")
			}
			if strings.Contains(errorMsg, "该用户已属于某个家庭，无法被邀请") {
				return errors.New("该用户已属于某个家庭，无法被邀请")
			}
		}
		return result.Error
	}
	return nil
}

// SetFamilyBudget 调用存储过程设置家庭预算
func SetFamilyBudget(userID int, budgetTime string, amount float64) error {
	// 调用存储过程 sp_set_family_budget
	result := DB.Exec("CALL sp_set_family_budget(?, ?, ?)", userID, budgetTime, amount)
	if result.Error != nil {
		// 检查是否是MySQL 45000异常（业务逻辑错误）
		errorMsg := result.Error.Error()
		if strings.Contains(errorMsg, "Error 1644") {
			if strings.Contains(errorMsg, "只有家庭管理员才能设置家庭预算") {
				return errors.New("只有家庭管理员才能设置家庭预算")
			}
		}
		return result.Error
	}
	return nil
}
