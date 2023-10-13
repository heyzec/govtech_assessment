package dataaccess

import (
    "gorm.io/gorm"

    "github.com/heyzec/govtech-assignment/internal/models"
)

func CreateTeacher(db *gorm.DB, teacher *models.Teacher) (*models.Teacher, error) {
    result := db.Create(teacher)

    if err := result.Error; err != nil {
        return nil, err
    }
    return teacher, nil
}

func ListTeachers(db *gorm.DB) ([]models.Teacher, error) {
    var teacherList []models.Teacher
    result := db.Model(&models.Teacher{}).Find(&teacherList)

    if err := result.Error; err != nil {
        return nil, err
    }
    return teacherList, nil
}
