package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"` // Hashed later
	RoleID   uint   `json:"role_id"`
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name     string `json:"name"`
}
