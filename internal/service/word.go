package service

import (
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/repository"
)

type WordService struct {
	repository *repository.WordRepository
	userRepository *repository.UserRepository
	languageRepository *repository.LanguageRepository
}

func NewWordService(repo *repository.WordRepository, userRepo *repository.UserRepository, languageRepo *repository.LanguageRepository) *WordService {
	return &WordService{
		repository: repo,
		userRepository: userRepo,
		languageRepository: languageRepo,
	}
}

func (w *WordService) FindByID(id uint) (*models.Word, error) {
	return w.repository.FindByID(id)
}

func (w *WordService) GetAll(limit int, username string, languageCode string, different bool) ([]models.Word, error) {
	var userID uint = 0
	var languageID uint = 0
	if username != "" {
		user, err := w.userRepository.FindByUsername(username)
		if err != nil {
			return []models.Word{}, err
		}
		userID = user.ID
	}
	if languageCode != "" {
		language, err := w.languageRepository.FindByCode(languageCode)
		if err != nil {
			return []models.Word{}, err
		}
		languageID = language.ID
	}
	return w.repository.GetAll(limit, userID, languageID, different)
}

func (w *WordService) Save(word *models.Word) (*models.Word, error) {
	lang, err := w.languageRepository.FindByCode(word.Language.Code)
	if err != nil {
		return nil, err
	}
	word.Language.ID = lang.ID
	return w.repository.Save(word)
}

func (w *WordService) Delete(userID uint, wordID uint) error {
	return w.repository.Delete(userID, wordID)
}

func (w *WordService) Update(word *models.Word) (*models.Word, error) {
	if word.Language.Code != "" {
		lang, err := w.languageRepository.FindByCode(word.Language.Code)
		if err != nil {
			return nil, err
		}
		word.Language.ID = lang.ID
	}
	return w.repository.Update(word)
}