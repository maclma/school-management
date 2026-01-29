package services

import (
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/repositories"
)

type ClassService struct {
	repo *repositories.ClassRepository
}

func NewClassService(repo *repositories.ClassRepository) *ClassService {
	return &ClassService{repo: repo}
}

func (s *ClassService) CreateClass(class *models.Class) error {
	return s.repo.Create(class)
}

func (s *ClassService) GetClasses() ([]models.Class, error) {
	return s.repo.GetAll()
}

func (s *ClassService) GetClassByID(id uint) (*models.Class, error) {
	return s.repo.GetByID(id)
}

func (s *ClassService) UpdateClass(class *models.Class) error {
	return s.repo.Update(class)
}

func (s *ClassService) DeleteClass(id uint) error {
	return s.repo.Delete(id)
}
