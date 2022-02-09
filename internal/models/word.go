package models

type Word struct {
	Name string `json:"name"`
	Mean string `json:"mean"`
	Story string `json:"story"`
	UserID uint `json:"user_id"`
	User *User `json:"user"`
}