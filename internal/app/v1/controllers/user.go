package controllers

import (
	"github.com/ahmetberke/wooker-api/internal/service"
)

type User struct {
	Service *service.UserService
}
