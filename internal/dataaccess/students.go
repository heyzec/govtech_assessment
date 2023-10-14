package dataaccess

import (
    "errors"
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

func FindStudentsByEmail(db *gorm.DB, emails []string) ([]models.Student, error) {
    var studentList []models.Student
    for _, email := range emails {
        var student models.Student
        err := db.Model(models.Student{}).
            Preload("Teachers").
            Where("email = ?", email).
            First(&student).
            Error
        if err != nil {
            return nil, errors.New("Student not found")
        }
        studentList = append(studentList, student)
    }

    return studentList, nil
}
