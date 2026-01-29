package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name    string `gorm:"not null;unique"`
	ClassID uint   `gorm:"not null"`
}
