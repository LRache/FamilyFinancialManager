package model

type Family struct {
	FamilyID    int     `gorm:"column:familyid;primaryKey;autoIncrement" json:"familyid"`
	FamilyName  string  `gorm:"column:familyname;not null" json:"familyname"`
	MonthBudget float64 `gorm:"column:monthbudget;default:0.00" json:"monthbudget"`
}

func (Family) TableName() string {
	return "Family"
}
