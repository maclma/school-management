package services

import (
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) Update(user *models.User) error {
	return s.repo.Update(user)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}
