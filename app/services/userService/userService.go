package userService

import (
	"BBS/app/models"
	"BBS/config/database"
)

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("user = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func Register(user models.User) error {
	result := database.DB.Create(&user)
	return result.Error
}