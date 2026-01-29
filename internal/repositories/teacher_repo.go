package repositories

import (
	"github.com/maclma/school-management/internal/models"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	DB *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) *TeacherRepository {
	return &TeacherRepository{DB: db}
}

func (r *TeacherRepository) Create(teacher *models.Teacher) error {
	return r.DB.Create(teacher).Error
}

func (r *TeacherRepository) GetAll() ([]models.Teacher, error) {
	var teachers []models.Teacher
	err := r.DB.Find(&teachers).Error
	return teachers, err
}

func (r *TeacherRepository) GetByID(id uint) (*models.Teacher, error) {
	var teacher models.Teacher
	err := r.DB.First(&teacher, id).Error
	if err != nil {
		return nil, err
	}
	return &teacher, err
}
func (r *TeacherRepository) Update(teacher *models.Teacher) error {
	return r.DB.Save(teacher).Error
}
func (r *TeacherRepository) GetPaginated(page, limit int) ([]models.Teacher, int64, error) {
	var teachers []models.Teacher
	var total int64

	offset := (page - 1) * limit

	if err := r.DB.Model(&models.Teacher{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Offset(offset).Limit(limit).Find(&teachers).Error; err != nil {
		return nil, 0, err
	}

	return teachers, total, nil
}

func (r *TeacherRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Teacher{}, id).Error
}
