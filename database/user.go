package database

import (
	"user-management/models"

	"github.com/gin-gonic/gin"
)

func GetUser(param string) []models.User {
	var user []models.User
	if param != "" {
		DB.Where("id = ?", param).Find(&user)
	} else {
		DB.Find(&user)
	}
	return user
}

func CreateUser(user *models.User) (*models.User, error) {
	if err := DB.Create(&user); err.Error != nil {
		return nil, err.Error
	}
	return user, nil
}

func UpdateUser(context *gin.Context, userID string, user *models.User) error {
	var userData models.User
	if err := DB.Where("id = ?", userID).Find(&userData).Error; err != nil {
		return err
	}
	DB.Model(userData).Updates(user)
	return nil
}

func DeleteUser(userID string) error {
	var userData models.User
	if err := DB.Where("id = ?", userID).Find(&userData).Error; err != nil {
		return err
	}
	DB.Delete(&userData)
	return nil
}
