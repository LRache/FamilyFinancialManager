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
	var user model.User
	err := DB.Where("UserName = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
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
