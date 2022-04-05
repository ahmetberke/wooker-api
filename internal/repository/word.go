package repository

import (
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/jackc/pgconn"
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
	err = w.db.Preload("Language").Preload("User").Find(&word).Error
	if err != nil {
		return nil, err
	}
	return word, nil
}

func (w *WordRepository) GetAll(limit int, userID uint, languageID uint, different bool) ([]models.Word, error) {
	var words []models.Word
	tx := w.db.Limit(limit)
	if different {
		tx.Distinct("name")
	}
	if userID != 0 {
		tx.Where("user_id = ?", userID)
	}
	if languageID != 0 {
		tx.Where("language_id = ?", languageID)
	}
	tx.Preload("User").Preload("Language").Find(&words)
	return words, nil
}

func (w *WordRepository) Save(word *models.Word) (*models.Word, error) {
	err := w.db.Where("name = ?", word.Name).Where("user_id = ?", word.UserID).Where("language_id = ?", word.Language.ID).First(&word).Error
	if err == nil {
		return nil, &pgconn.PgError{}
	}
	err = w.db.Create(&word).Error
	if err != nil {
		return nil, err
	}
	err = w.db.Preload("User").Preload("Language").Find(&word).Error
	if err != nil {
		return nil, err
	}
	return word, nil
}

func (w *WordRepository) Delete(userID uint, wordID uint) error {
	var word models.Word
	err := w.db.Where("user_id = ?", userID).Where("id = ?", wordID).Delete(&word).Error
	if err != nil {
		return err
	}
	return nil
}

func (w* WordRepository) Update(word *models.Word) (*models.Word, error) {
	err := w.db.Where("user_id = ?", word.UserID).Where("id = ?", word.ID).Updates(&word).Error
	if err != nil {
		return nil, err
	}
	return word, err
}
