package repository

import (
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func (s *SessionRepository) Get(token uuid.UUID) (*models.Session, error)  {
	var session *models.Session
	err := s.db.Where("token=?", token).First(&session).Error
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SessionRepository) Save(session *models.Session) (*models.Session, error)  {
	err := s.db.Save(&session).Error
	if err != nil {
		return nil, err
	}
	return session, err
}