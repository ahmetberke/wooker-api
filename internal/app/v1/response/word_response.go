package response

import "github.com/ahmetberke/wooker-api/internal/models"

type WordResponse struct {
	Response
	Word *models.WordDTO `json:"word"`
}

type WordsResponse struct {
	Response
	Words []*models.WordDTO `json:"words"`
}