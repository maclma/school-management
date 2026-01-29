package services

import (
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/repositories"
)

type AttendanceService struct {
	repo *repositories.AttendanceRepository
}

func NewAttendanceService(repo *repositories.AttendanceRepository) *AttendanceService {
	return &AttendanceService{repo: repo}
}
func (s *AttendanceService) GetPaginatedAttendance(page, limit int) ([]models.Attendance, int64, error) {
	return s.repo.GetPaginated(page, limit)
}

func (s *AttendanceService) Create(attendance *models.Attendance) error {
	return s.repo.Create(attendance)
}
func (s *AttendanceService) MarkAttendance(attendance *models.Attendance) error {
	return s.repo.Create(attendance)
}

func (s *AttendanceService) GetAllAttendance() ([]models.Attendance, error) {
	return s.repo.GetAll()
}

func (s *AttendanceService) GetAttendanceByStudent(studentID uint) ([]models.Attendance, error) {
	return s.repo.GetByStudent(studentID)
}

func (s *AttendanceService) GetAttendanceByClassAndDate(classID uint, date string) ([]models.Attendance, error) {
	return s.repo.GetByClassAndDate(classID, date)
}
