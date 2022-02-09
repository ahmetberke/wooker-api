package service

import (
	"fmt"
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repository: repo}
}

func (u *UserService) Save(user *models.User) (*models.User, error) {
	return u.repository.Save(user)
}

func (u *UserService) FindByUsername(username string) (*models.User, error)  {
	return u.repository.FindByUsername(username)
}

func (u *UserService) MultipleFindByUsername(username string) ([]models.User, error)  {
	return u.repository.MultiFindByUsername(username)
}

func (u *UserService) FindByID(id uint) (*models.User, error)  {
	return u.repository.FindByID(id)
}

func (u *UserService) FindByGoogleID(id string) (*models.User, error)  {
	return u.repository.FindByGoogleID(id)
}

func (u *UserService) GetAll(limit int) ([]models.User, error)  {
	return u.repository.GetAll(limit)
}

func (u *UserService) UpdateByUsername(username string, user *models.User) (*models.User, error) {
	if username == user.Username {
		return nil, fmt.Errorf("the username is already the same")
	}

	user.ID = 0
	user.Email = ""
	user.EmailVerified = false
	user.IsAdmin = false
	user.Picture = ""

	return u.repository.UpdateByUsername(username, user)
}

func (u *UserService) Delete(id uint) error  {
	return u.repository.Delete(id)
}

func (u *UserService) Migrate() error {
	return u.repository.Migrate()
}