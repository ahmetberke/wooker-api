package repository

import (
	"github.com/ahmetberke/wooker-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Save(user *models.User) (*models.User, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user *models.User
	result := u.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepository) MultiFindByUsername(username string) ([]*models.User, error)  {
	var users []*models.User
	result := u.db.Find(users,&models.User{Username: username})
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (u *UserRepository) FindByID(id uint) (*models.User, error) {
	var user *models.User
	result := u.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepository) FindByGoogleID(id string) (*models.User, error) {
	var user *models.User
	result := u.db.Where("google_id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepository) Update(user *models.User) (*models.User, error)  {
	result := u.db.Model(&user).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepository) Delete(id uint) error {
	result := u.db.Delete(&models.User{ID:id})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) Migrate() error {
	return u.db.AutoMigrate(&models.User{})
}