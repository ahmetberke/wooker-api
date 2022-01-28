package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID string `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"primaryKey"`
	Email string `json:"email" gorm:"primaryKey"`
	EmailVerified string `json:"email_verified"`
	Picture string `json:"picture"`
	CreatedDate string `json:"created_date"`
	LastUpdatedDate string `json:"last_updated_date"`
}

type  UserDTO struct {
	Username string `json:"username"`
	Email string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Picture string `json:"picture"`
}
