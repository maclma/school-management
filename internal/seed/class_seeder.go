package seed

import (
	"github.com/maclma/school-management/internal/models"
	"gorm.io/gorm"
)

func SeedClasses(db *gorm.DB) error {
	classes := []models.Class{
		{Name: "Class 1"},
		{Name: "Class 2"},
		{Name: "Class 3"},
	}

	for _, class := range classes {
		var existing models.Class
		if err := db.Where("name = ?", class.Name).First(&existing).Error; err == nil {
			continue
		}
		if err := db.Create(&class).Error; err != nil {
			return err
		}
	}
	return nil
}
