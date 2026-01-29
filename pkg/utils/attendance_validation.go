package utils

import (
	"time"

	"github.com/maclma/school-management/internal/models"
)

func ValidateAttendance(a *models.Attendance) error {
	if a.StudentID == 0 {
		return BadRequest("student_id is required")
	}

	if a.ClassID == 0 {
		return BadRequest("class_id is required")
	}

	if a.Date.IsZero() {
		return BadRequest("date is required")
	}

	// Validate date format YYYY-MM-DD
	if _, err := time.Parse("2006-01-02", a.Date.Format("2006-01-02")); err != nil {
		return BadRequest("date must be in YYYY-MM-DD format")
	}
	if a.Status == "" {
		return BadRequest("status is required")
	}

	if a.Status != "present" && a.Status != "absent" {
		return BadRequest("status must be 'present' or 'absent'")
	}
	if a.Date.After(time.Now()) {
		return BadRequest("date cannot be in the future")
	}
	return nil
}
