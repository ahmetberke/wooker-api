package response

import "github.com/ahmetberke/wooker-api/internal/models"

type LanguageResponse struct {
	Response
	Language *models.Language `json:"language"`
}

type LanguagesResponse struct {
	Response
	Languages []models.Language `json:"languages"`
}