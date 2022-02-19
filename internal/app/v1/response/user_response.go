package response

import "github.com/ahmetberke/wooker-api/internal/models"

type UserResponse struct {
	Response
	User *models.UserDTO `json:"user"`
	LoggedIn bool `json:"logged_in"`
}

type UsersResponse struct {
	Response
	Users []models.UserDTO `json:"users"`
}

type AuthResponse struct {
	UserResponse
	Token string `json:"token"`
}