package models

import (
	"github.com/heyzec/govtech-assignment/internal/helpers/emailutils"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Email     string    `gorm:"not null"`
	Teachers  []Teacher `gorm:"many2many:students_teachers;"`
	Suspended bool      `gorm:"not null;"`
}

func (student *Student) Validate() error {
	err := emailutils.ValidateEmailValid(student.Email)
	if err != nil {
		return err
	}
	return nil
}

// ===== GORM Hooks =====
func (student *Student) BeforeCreate(db *gorm.DB) error {
	if err := student.Validate(); err != nil {
		return err
	}
	return nil
}
