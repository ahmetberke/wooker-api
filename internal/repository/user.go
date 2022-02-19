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
	err := u.db.Create(&user).Preload("User").Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user *models.User
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) MultiFindByUsername(username string) ([]models.User, error)  {
	var users []models.User
	err := u.db.Where("username = ?", username).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) FindByID(id uint) (*models.User, error) {
	var user *models.User
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) FindByGoogleID(id string) (*models.User, error) {
	var user *models.User
	err := u.db.Where("google_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetAll(limit int, search string) ([]models.User, error)  {
	var users []models.User
	tx := u.db.Limit(limit)
	if search != "" {
		tx.Where("username LIKE ?", search+"%")
	}
	err := tx.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (u *UserRepository) UpdateByUsername(username string, user *models.User) (*models.User, error)  {
	err := u.db.Model(&user).Where("username = ?", username).Updates(&user).Preload("User").Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) Delete(id uint) error {
	err := u.db.Delete(&models.User{ID:id}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Migrate() error {
	return u.db.AutoMigrate(&models.User{})
}