package service

import (
	"backend/api/response"
	"backend/internal/repository"
	"net/http"
	"strings"

	"github.com/wonderivan/logger"
)

// CreateFamily 创建家庭
func CreateFamily(userID int, familyName string) *Result[response.CreateFamily] {
	familyID, err := repository.CreateFamily(userID, familyName)
	if err != nil {
		logger.Warn("Create family error:", err.Error())

		// 检查是否是业务逻辑错误
		if strings.Contains(err.Error(), "该用户已属于某个家庭") {
			return ResultFailed[response.CreateFamily](http.StatusConflict, err.Error())
		}

		return ResultFailed[response.CreateFamily](http.StatusInternalServerError, "Internal server error")
	}

	return ResultOK(response.CreateFamily{
		FamilyID:   familyID,
		FamilyName: familyName,
	})
}

// InviteUserToFamily 邀请用户加入家庭
func InviteUserToFamily(inviterID int, username string) *Result[response.InviteUser] {
	// 首先根据用户名查找用户ID
	inviteeUser, err := repository.GetUserByUsername(username)
	if err != nil {
		logger.Warn("Get user error:", err.Error())
		return ResultFailed[response.InviteUser](http.StatusInternalServerError, "Internal server error")
	}

	if inviteeUser == nil {
		return ResultFailed[response.InviteUser](http.StatusNotFound, "用户不存在")
	}

	err = repository.InviteUserToFamily(inviterID, inviteeUser.UserID)
	if err != nil {
		logger.Warn("Invite user error:", err.Error())

		// 检查是否是业务逻辑错误
		if strings.Contains(err.Error(), "只有家庭管理员才能邀请") {
			return ResultFailed[response.InviteUser](http.StatusForbidden, err.Error())
		}
		if strings.Contains(err.Error(), "该用户已属于某个家庭") {
			return ResultFailed[response.InviteUser](http.StatusConflict, err.Error())
		}

		return ResultFailed[response.InviteUser](http.StatusInternalServerError, "Internal server error")
	}

	return ResultOK(response.InviteUser{
		Username: username,
		Status:   "邀请成功",
	})
}

// GetFamilyMembers 获取家庭成员列表
func GetFamilyMembers(userID int) *Result[response.FamilyMembers] {
	// 首先获取用户信息以确定家庭ID
	user, err := repository.GetUserByID(userID)
	if err != nil {
		logger.Warn("Get user error:", err.Error())
		return ResultFailed[response.FamilyMembers](http.StatusInternalServerError, "Internal server error")
	}

	if user == nil {
		return ResultFailed[response.FamilyMembers](http.StatusNotFound, "用户不存在")
	}

	if user.FamilyID == nil {
		return ResultFailed[response.FamilyMembers](http.StatusBadRequest, "用户未加入任何家庭")
	}

	members, err := repository.GetFamilyMembers(*user.FamilyID)
	if err != nil {
		logger.Warn("Get family members error:", err.Error())
		return ResultFailed[response.FamilyMembers](http.StatusInternalServerError, "Internal server error")
	}

	// 转换为响应格式
	var familyMembers []response.FamilyMember
	for _, member := range members {
		familyMembers = append(familyMembers, response.ConvertUserToFamilyMember(member))
	}

	return ResultOK(response.FamilyMembers{
		Members: familyMembers,
	})
}

// SetFamilyBudget 设置家庭预算
func SetFamilyBudget(userID int, budgetTime string, amount float64) *Result[string] {
	err := repository.SetFamilyBudget(userID, budgetTime, amount)
	if err != nil {
		logger.Warn("Set family budget error:", err.Error())

		// 检查是否是业务逻辑错误
		if strings.Contains(err.Error(), "只有家庭管理员才能设置") {
			return ResultFailed[string](http.StatusForbidden, err.Error())
		}

		return ResultFailed[string](http.StatusInternalServerError, "Internal server error")
	}

	return ResultOK("预算设置成功")
}

// GetFamilyInfo 获取家庭信息
func GetFamilyInfo(userID int) *Result[response.FamilyInfo] {
	// 首先获取用户信息以确定家庭ID
	user, err := repository.GetUserByID(userID)
	if err != nil {
		logger.Warn("Get user error:", err.Error())
		return ResultFailed[response.FamilyInfo](http.StatusInternalServerError, "Internal server error")
	}

	if user == nil {
		return ResultFailed[response.FamilyInfo](http.StatusNotFound, "用户不存在")
	}

	if user.FamilyID == nil {
		return ResultFailed[response.FamilyInfo](http.StatusBadRequest, "用户未加入任何家庭")
	}

	// 获取家庭信息
	family, err := repository.GetFamilyByID(*user.FamilyID)
	if err != nil {
		logger.Warn("Get family error:", err.Error())
		return ResultFailed[response.FamilyInfo](http.StatusInternalServerError, "Internal server error")
	}

	if family == nil {
		return ResultFailed[response.FamilyInfo](http.StatusNotFound, "家庭不存在")
	}

	// 获取成员数量
	members, err := repository.GetFamilyMembers(*user.FamilyID)
	if err != nil {
		logger.Warn("Get family members error:", err.Error())
		return ResultFailed[response.FamilyInfo](http.StatusInternalServerError, "Internal server error")
	}

	return ResultOK(response.FamilyInfo{
		FamilyID:    family.FamilyID,
		FamilyName:  family.FamilyName,
		MonthBudget: family.MonthBudget,
		MemberCount: len(members),
	})
}
