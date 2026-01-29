package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name     string    `gorm:"not null;unique"`
	Students []Student `gorm:"foreignKey:ClassID"`
	Subjects []Subject `gorm:"foreignKey:ClassID"`
}
