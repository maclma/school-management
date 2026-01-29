package repositories

import (
	"github.com/maclma/school-management/internal/models"
	"gorm.io/gorm"
)

type ClassRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{db: db}
}

func (r *ClassRepository) Create(class *models.Class) error {
	return r.db.Create(class).Error
}

func (r *ClassRepository) GetAll() ([]models.Class, error) {
	var classes []models.Class
	err := r.db.
		Preload("Students").
		Preload("Subjects").
		Find(&classes).Error
	return classes, err
}

func (r *ClassRepository) GetByID(id uint) (*models.Class, error) {
	var class models.Class
	err := r.db.
		Preload("Students").
		Preload("Subjects").
		First(&class, id).Error
	if err != nil {
		return nil, err
	}
	return &class, nil
}

func (r *ClassRepository) Update(class *models.Class) error {
	return r.db.Save(class).Error
}

func (r *ClassRepository) Delete(id uint) error {
	return r.db.Delete(&models.Class{}, id).Error
}
