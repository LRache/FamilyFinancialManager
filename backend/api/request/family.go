package request

type CreateFamily struct {
	FamilyName string `form:"familyname" json:"familyname" binding:"required"`
}

type InviteUser struct {
	Username string `form:"username" json:"username" binding:"required"`
}

type SetBudget struct {
	Amount    float64 `form:"amount" json:"amount" binding:"required"`
	StartDate string  `form:"start_date" json:"start_date" binding:"required"` // 格式: YYYY-MM-DD
	Category  string  `form:"category" json:"category"`
	Note      string  `form:"note" json:"note"`
}
