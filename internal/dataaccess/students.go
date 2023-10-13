package dataaccess

import (
    "gorm.io/gorm"

    "github.com/heyzec/govtech-assignment/internal/models"
)

func CreateStudent(db *gorm.DB, student *models.Student) (*models.Student, error) {
    result := db.Create(student)

    if err := result.Error; err != nil {
        return nil, err
    }
    return student, nil
}

func ListStudents(db *gorm.DB) ([]models.Student, error) {
    var studentList []models.Student
    result := db.Model(&models.Student{}).Find(&studentList)

    if err := result.Error; err != nil {
        return nil, err
    }
    return studentList, nil
}
