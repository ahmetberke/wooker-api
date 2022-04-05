package models

import "time"

type Word struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Mean string `json:"mean"`
	Story string `json:"story"`
	UserID uint `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:UserID"`
	LanguageID uint `json:"language_id"`
	Language Language `json:"language" gorm:"foreignKey:LanguageID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WordDTO struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Mean string `json:"mean"`
	Story string `json:"story"`
	User UserDTO `json:"user"`
	Language Language `json:"language"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToWord(wordDTO *WordDTO) *Word  {
	return &Word{
		ID:         wordDTO.ID,
		Name:       wordDTO.Name,
		Mean:       wordDTO.Mean,
		Story:      wordDTO.Story,
		User: 		*ToUser(&wordDTO.User),
		Language: 	wordDTO.Language,
		CreatedAt:  wordDTO.CreatedAt,
		UpdatedAt:  wordDTO.UpdatedAt,
	}
}

func ToWordDTO(word *Word) *WordDTO {
	return &WordDTO{
		ID:        word.ID,
		Name:      word.Name,
		Mean:      word.Mean,
		Story:     word.Story,
		User: 	   *ToUserDTO(&word.User),
		Language:  word.Language,
		CreatedAt: word.CreatedAt,
		UpdatedAt: word.UpdatedAt,
	}
}