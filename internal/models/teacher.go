package models

import (
    "gorm.io/gorm"
)

type Teacher struct {
    gorm.Model
    Email string `gorm:"not null"`
    Students []Student `gorm:"many2many:students_teachers;"`
}
