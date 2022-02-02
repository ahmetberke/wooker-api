package service

import "github.com/ahmetberke/wooker-api/internal/repository"

type SessionService struct {
	SessionRepository *repository.SessionRepository
}