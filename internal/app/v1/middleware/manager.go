package middleware

import "github.com/ahmetberke/wooker-api/internal/service"

type Manager struct {
	UserService *service.UserService
}

func NewManager(userService *service.UserService) *Manager {
	return &Manager{
		UserService: userService,
	}
}