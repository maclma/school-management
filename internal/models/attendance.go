package models

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	StudentID uint      `gorm:"not null;index:idx_attendance_unique,unique"`
	ClassID   uint      `gorm:"not null;index:idx_attendance_unique,unique"`
	Date      time.Time `gorm:"not null;index:idx_attendance_unique,unique"`
	Status    string    `gorm:"not null"` // Present / Absent
}
