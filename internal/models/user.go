package models

import (
	"time"
)

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	GoogleID string `json:"google_id" gorm:"unique"`
	Username string `json:"username" gorm:"unique"`
	IsAdmin bool `json:"is_admin"`
	Email string `json:"email" gorm:"primaryKey"`
	EmailVerified bool `json:"email_verified"`
	Picture string `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type  UserDTO struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	IsAdmin bool `json:"is_admin"`
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