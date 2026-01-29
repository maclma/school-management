package repositories

import (
	"github.com/maclma/school-management/internal/models"
	"gorm.io/gorm"
)

type SubjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) *SubjectRepository {
	return &SubjectRepository{db: db}
}

func (r *SubjectRepository) Create(subject *models.Subject) error {
	return r.db.Create(subject).Error
}

func (r *SubjectRepository) GetAll() ([]models.Subject, error) {
	var subjects []models.Subject
	err := r.db.Find(&subjects).Error
	return subjects, err
}

func (r *SubjectRepository) GetByID(id uint) (*models.Subject, error) {
	var subject models.Subject
	err := r.db.First(&subject, id).Error
	return &subject, err
}

func (r *SubjectRepository) GetByClassID(classID uint) ([]models.Subject, error) {
	var subjects []models.Subject
	err := r.db.Where("class_id = ?", classID).Find(&subjects).Error
	return subjects, err
}
func (r *SubjectRepository) GetPaginated(page, limit int) ([]models.Subject, int64, error) {
	var subjects []models.Subject
	var total int64
	offset := (page - 1) * limit

	if err := r.db.Model(&models.Subject{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Offset(offset).Limit(limit).Find(&subjects).Error; err != nil {
		return nil, 0, err
	}
	return subjects, total, nil
}

func (r *SubjectRepository) Update(subject *models.Subject) error {
	return r.db.Save(subject).Error
}

func (r *SubjectRepository) Delete(id uint) error {
	return r.db.Delete(&models.Subject{}, id).Error
}
