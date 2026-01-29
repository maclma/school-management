package seed

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint
	Name     string
	Email    string `gorm:"unique"`
	Password string
	RoleID   uint
}

func SeedAdminUser(db *gorm.DB) error {
	var adminRole Role
	if err := db.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		return err
	}

	var existing User
	if err := db.Where("email = ?", "admin@school.com").First(&existing).Error; err == nil {
		return nil // admin already exists
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := User{
		Name:     "System Admin",
		Email:    "admin@school.com",
		Password: string(hashed),
		RoleID:   adminRole.ID,
	}

	return db.Create(&admin).Error
}
