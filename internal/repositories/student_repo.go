package repositories

import (
	"github.com/maclma/school-management/internal/models"
	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db: db}
}
func (r *StudentRepository) GetAll() ([]models.Student, error) {
	var students []models.Student
	err := r.db.Preload("Class").Find(&students).Error
	return students, err
}

// ✅ New paginated method (SAFE)
func (r *StudentRepository) GetPaginated(offset, limit int) ([]models.Student, int64, error) {
	var students []models.Student
	var total int64

	if err := r.db.Model(&models.Student{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset(offset).Limit(limit).Find(&students).Error; err != nil {
		return nil, 0, err
	}

	return students, total, nil
}

func (r *StudentRepository) Create(student *models.Student) error {
	return r.db.Create(student).Error
}

func (r *StudentRepository) GetByID(id uint) (*models.Student, error) {
	var student models.Student
	err := r.db.First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) Update(student *models.Student) error {
	return r.db.Save(student).Error
}

func (r *StudentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Student{}, id).Error
}
