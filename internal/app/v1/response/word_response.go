package response

import "github.com/ahmetberke/wooker-api/internal/models"

type WordResponse struct {
	Response
	Word *models.WordDTO `json:"user"`
}

type WordsResponse struct {
	Response
	Words []*models.WordDTO `json:"user"`
}