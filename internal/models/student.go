package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	ClassID   uint   `json:"class_id" gorm:"not null"`
	Class     Class  `json:"class" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
