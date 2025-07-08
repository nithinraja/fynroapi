package user

import (
	"ai-financial-api/config"
	"ai-financial-api/models"
)

func GetUserByUUID(uuid string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
