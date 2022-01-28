package service

import (
	"github.com/ahmetberke/wooker-api/internal/models"
	"gorm.io/gorm"
)

type UserService struct {
	Database *gorm.DB
}

func (us *UserService) Create(user *models.User) (*models.User, error) {
	result := us.Database.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (us *UserService) GetByUsername(username string) (*models.User, error) {
	var user *models.User
	result := us.Database.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (us *UserService) Update(user *models.User) (*models.User, error)  {
	result := us.Database.Model(&user).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}