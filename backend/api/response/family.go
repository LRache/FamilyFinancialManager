package response

import "backend/internal/model"

type CreateFamily struct {
	FamilyID   int    `json:"familyid"`
	FamilyName string `json:"familyname"`
}

type InviteUser struct {
	Username string `json:"username"`
	Status   string `json:"status"`
}

type FamilyMembers struct {
	Members []FamilyMember `json:"members"`
}

type FamilyMember struct {
	Username string  `json:"username"`
	Email    *string `json:"email"`
	Role     int     `json:"role"` // 0=家庭成员，1=家庭管理员
}

type FamilyInfo struct {
	FamilyID    int     `json:"familyid"`
	FamilyName  string  `json:"familyname"`
	MonthBudget float64 `json:"monthbudget"`
	MemberCount int     `json:"member_count"`
}

// ConvertUserToFamilyMember 将User model转换为FamilyMember响应
func ConvertUserToFamilyMember(user model.User) FamilyMember {
	return FamilyMember{
		Username: user.UserName,
		Email:    user.Email,
		Role:     user.Role,
	}
}
