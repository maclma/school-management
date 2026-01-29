package seed

import "gorm.io/gorm"

type Role struct {
	ID   uint
	Name string `gorm:"unique"`
}

func SeedRoles(db *gorm.DB) error {
	roles := []Role{
		{Name: "admin"},
		{Name: "teacher"},
		{Name: "student"},
	}

	for _, role := range roles {
		var existing Role
		if err := db.Where("name = ?", role.Name).First(&existing).Error; err == nil {
			continue // already exists
		}
		if err := db.Create(&role).Error; err != nil {
			return err
		}
	}
	return nil
}
