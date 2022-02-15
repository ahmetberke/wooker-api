package repository

import (
	"github.com/ahmetberke/wooker-api/internal/models"
	"gorm.io/gorm"
)

type LanguageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) *LanguageRepository {
	return &LanguageRepository{
		db: db,
	}
}

func (l *LanguageRepository) FindByID(id uint) (*models.Language, error) {
	var language *models.Language
	err := l.db.First(&language, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return language, nil
}

func (l *LanguageRepository) FindByName(name string) (*models.Language, error) {
	var language *models.Language
	err := l.db.First(&language, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	return language, nil
}

func (l *LanguageRepository) FindByCode(code string) (*models.Language, error)  {
	var language *models.Language
	err := l.db.First(&language, "code = ?", code).Error
	if err != nil {
		return nil, err
	}
	return language, err
}