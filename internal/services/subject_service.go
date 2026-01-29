package services

import (
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/repositories"
)

type SubjectService struct {
	repo *repositories.SubjectRepository
}

func NewSubjectService(repo *repositories.SubjectRepository) *SubjectService {
	return &SubjectService{repo: repo}
}

/* ===== CREATE ===== */

func (s *SubjectService) CreateSubject(subject *models.Subject) error {
	return s.repo.Create(subject)
}

/* ===== READ ===== */

func (s *SubjectService) GetSubjects() ([]models.Subject, error) {
	return s.repo.GetAll()
}

func (s *SubjectService) GetSubjectByID(id uint) (*models.Subject, error) {
	return s.repo.GetByID(id)
}

func (s *SubjectService) GetSubjectsByClass(classID uint) ([]models.Subject, error) {
	return s.repo.GetByClassID(classID)
}

/* ===== UPDATE ===== */

func (s *SubjectService) UpdateSubject(subject *models.Subject) error {
	return s.repo.Update(subject)
}
func (s *SubjectService) GetPaginatedSubjects(page, limit int) ([]models.Subject, int64, error) {
	return s.repo.GetPaginated(page, limit)
}

/* ===== DELETE ===== */

func (s *SubjectService) DeleteSubject(id uint) error {
	return s.repo.Delete(id)
}
