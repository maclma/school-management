package repositories

import (
	"github.com/maclma/school-management/internal/models"
	"gorm.io/gorm"
)

type AttendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{db: db}
}

func (r *AttendanceRepository) Create(attendance *models.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *AttendanceRepository) GetAll() ([]models.Attendance, error) {
	var attendance []models.Attendance
	err := r.db.
		Preload("Student").
		Preload("Class").
		Find(&attendance).Error
	return attendance, err
}

func (r *AttendanceRepository) GetByStudent(studentID uint) ([]models.Attendance, error) {
	var attendance []models.Attendance
	err := r.db.
		Where("student_id = ?", studentID).
		Find(&attendance).Error
	return attendance, err
}
func (r *AttendanceRepository) GetPaginated(page, limit int) ([]models.Attendance, int64, error) {
	var attendance []models.Attendance
	var total int64
	offset := (page - 1) * limit
	if err := r.db.Model(&models.Attendance{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Offset(offset).Limit(limit).
		Preload("Student").
		Preload("Class").
		Find(&attendance).Error; err != nil {
		return nil, 0, err
	}
	return attendance, total, nil
}

func (r *AttendanceRepository) GetByClassAndDate(classID uint, date string) ([]models.Attendance, error) {
	var attendance []models.Attendance
	err := r.db.
		Where("class_id = ? AND date = ?", classID, date).
		Find(&attendance).Error
	return attendance, err
}
func (r *AttendanceRepository) MarkAttendance(attendances []models.Attendance) error {
	tx := r.db.Begin()
	for _, attendance := range attendances {
		if err := tx.Create(&attendance).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
