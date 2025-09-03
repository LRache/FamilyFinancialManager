package repository

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

func CreateUser(username string, password string) (int, error) {
	result := DB.Exec("CALL sp_register_api(?, ?, @userid)", username, password)
	if result.Error != nil {
		return 0, result.Error
	}

	var userID int
	err := DB.Raw("SELECT @userid").Scan(&userID).Error
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func UserExists(username string) (bool, error) {
	var count int64
	query := DB.Table("Users")
	err := query.Where("UserName = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func UserLogin(username string, password string) (int, error) {
	var user model.User
	err := DB.Where("UserName = ? AND Password = ?", username, password).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return -1, nil
		}
		return -1, err
	}
	return user.UserID, nil
}

// GetUserByUsername 根据用户名获取用户信息
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByID 根据用户ID获取用户信息
func GetUserByID(userID int) (*model.User, error) {
	var user model.User
	err := DB.Where("userid = ?", userID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
