package models

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
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
