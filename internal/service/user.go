package service

import (
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserSercive(repo *repository.UserRepository) *UserService {
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

func (u *UserService) Update(user *models.User) (*models.User, error) {
	return u.repository.Update(user)
}

func (u *UserService) Delete(id uint) error  {
	return u.repository.Delete(id)
}

func (u *UserService) Migrate() error {
	return u.repository.Migrate()
}