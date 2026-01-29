package services

import (
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/repositories"
)

type TeacherService struct {
	Repo *repositories.TeacherRepository
}

func NewTeacherService(repo *repositories.TeacherRepository) *TeacherService {
	return &TeacherService{Repo: repo}
}

func (s *TeacherService) CreateTeacher(teacher *models.Teacher) error {
	return s.Repo.Create(teacher)
}
func (s *TeacherService) GetPaginatedTeachers(page, limit int) ([]models.Teacher, int64, error) {
	return s.Repo.GetPaginated(page, limit)
}

func (s *TeacherService) GetTeachers() ([]models.Teacher, error) {
	return s.Repo.GetAll()
}

func (s *TeacherService) GetTeacherByID(id uint) (*models.Teacher, error) {
	return s.Repo.GetByID(id)
}

func (s *TeacherService) UpdateTeacher(teacher *models.Teacher) error {
	return s.Repo.Update(teacher)
}

func (s *TeacherService) DeleteTeacher(id uint) error {
	return s.Repo.Delete(id)
}
