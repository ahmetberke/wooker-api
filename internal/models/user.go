package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey"`
	GoogleID string `json:"google_id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"primaryKey"`
	Email string `json:"email" gorm:"primaryKey"`
	EmailVerified bool `json:"email_verified"`
	Picture string `json:"picture"`
	CreatedDate time.Time `json:"created_date"`
	LastUpdatedDate time.Time `json:"last_updated_date"`
}

type  UserDTO struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	EmailVerified bool `json:"email_verified"`
	Picture string `json:"picture"`
}

func ToUser(userDTO *UserDTO) *User {
	return &User{
		ID: userDTO.ID,
		Username: userDTO.Username,
		Email: userDTO.Email,
		EmailVerified: userDTO.EmailVerified,
		Picture: userDTO.Picture,
	}
}

func ToUserDTO(user *User) *UserDTO {
	return &UserDTO{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		EmailVerified: user.EmailVerified,
		Picture: user.Picture,
	}
}