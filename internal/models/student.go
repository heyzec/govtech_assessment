package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Email    string    `gorm:"not null"`
	Teachers []Teacher `gorm:"many2many:students_teachers;"`
}
