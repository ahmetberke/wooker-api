package repository

import (
	"github.com/ahmetberke/wooker-api/internal/models"
	"gorm.io/gorm"
)

type WordRepository struct {
	db *gorm.DB
}

func NewWordRepository(db *gorm.DB) *WordRepository {
	return &WordRepository{
		db: db,
	}
}

func (w *WordRepository) FindByID(id uint) (*models.Word, error) {
	var word *models.Word
	err := w.db.Where("id = ?", id).First(&word).Error
	if err != nil {
		return nil, err
	}
	err = w.db.Preload("Language").Find(&word).Error
	if err != nil {
		return nil, err
	}
	err = w.db.Preload("User").Find(&word).Error
	if err != nil {
		return nil, err
	}
	return word, nil
}

func (w *WordRepository) Save(word *models.Word) (*models.Word, error) {
	err := w.db.Create(&word).Error
	if err != nil {
		return nil, err
	}
	return word, nil
}

func (w *WordRepository) Delete(wordId uint) error {
	err := w.db.Delete(&models.Word{ID: wordId}).Error
	if err != nil {
		return err
	}
	return nil
}