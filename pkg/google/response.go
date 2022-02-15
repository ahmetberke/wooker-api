package google

type UserResponse struct {
	ID string `json:"id"`
	Email string `json:"email"`
	VerifiedEmail bool `json:"verified_email"`
	Picture string `json:"picture"`
	Error *Error   `json:"error"`
}

type Error struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
}