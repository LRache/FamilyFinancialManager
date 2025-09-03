package model

type User struct {
	UserID   int     `gorm:"column:userid;primaryKey;autoIncrement" json:"userid"`
	UserName string  `gorm:"column:username;not null" json:"username"`
	Password string  `gorm:"column:password;not null" json:"-"`
	Email    *string `gorm:"column:email;unique" json:"email"`
	FamilyID *int    `gorm:"column:familyid" json:"familyid"`
	Role     int     `gorm:"column:role;not null;default:0" json:"role"` // 0=家庭成员，1=家庭管理员
}

func (User) TableName() string {
	return "Users"
}
