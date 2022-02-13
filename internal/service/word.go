package service

import (
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/repository"
)

type WordService struct {
	repository *repository.WordRepository
}

func NewWordService(repo *repository.WordRepository) *WordService {
	return &WordService{repository: repo}
}

func (w *WordService) FindByID(id uint) (*models.Word, error) {
	return w.repository.FindByID(id)
}

func (w *WordService) GetAll(limit int) ([]models.Word, error) {
	return w.repository.GetAll(limit)
}

func (w *WordService) Save(word *models.Word) (*models.Word, error) {
	return w.repository.Save(word)
}

func (w *WordService) Delete(wordId uint) error {
	return w.repository.Delete(wordId)
}