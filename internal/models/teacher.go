package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	SubjectID uint
}
