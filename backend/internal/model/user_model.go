package model

type User struct {
	UserID    int    `gorm:"column:UserID;primaryKey"`
	UserName  string `gorm:"column:UserName;unique"`
	Password  string `gorm:"column:Password"`
	Email     string `gorm:"column:Email"`
	FamilyID  int    `gorm:"column:FamilyID"`
	Token     string `gorm:"column:Token"`
	CreatedAt string `gorm:"column:CreatedAt"`
	UpdatedAt string `gorm:"column:UpdatedAt"`
}

func (User) TableName() string {
	return "Users"
}
