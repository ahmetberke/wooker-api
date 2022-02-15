package service

import (
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/repository"
)

type LanguageService struct {
	repository *repository.LanguageRepository
}

func NewLanguageService(repository *repository.LanguageRepository) *LanguageService  {
	return &LanguageService{
		repository: repository,
	}
}

func (l *LanguageService) FindByID(id uint) (*models.Language, error)  {
	return l.repository.FindByID(id)
}

func (l *LanguageService) FindByName(name string) (*models.Language, error)  {
	return l.repository.FindByName(name)
}

func (l *LanguageService) FindByCode(code string) (*models.Language, error)  {
	return l.repository.FindByCode(code)
}
