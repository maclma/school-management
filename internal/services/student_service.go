package services

import (
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/repositories"
)

type StudentService struct {
	repo *repositories.StudentRepository
}

func NewStudentService(repo *repositories.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}
func (s *StudentService) GetAll() ([]models.Student, error) {
	return s.repo.GetAll()
}

// ✅ Paginated version
func (s *StudentService) GetPaginated(page, limit int) ([]models.Student, int64, error) {
	offset := (page - 1) * limit
	return s.repo.GetPaginated(offset, limit)
}

func (s *StudentService) CreateStudent(student *models.Student) error {
	return s.repo.Create(student)
}

func (s *StudentService) GetStudentByID(id uint) (*models.Student, error) {
	return s.repo.GetByID(id)
}

func (s *StudentService) UpdateStudent(student *models.Student) error {
	return s.repo.Update(student)
}
func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.repo.GetAll()
}
func (s *StudentService) DeleteStudent(id uint) error {
	return s.repo.Delete(id)
}
