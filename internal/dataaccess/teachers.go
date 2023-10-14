package dataaccess

import (
	"errors"

	"github.com/heyzec/govtech-assignment/internal/models"
	"gorm.io/gorm"
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

func FindTeachersByEmail(db *gorm.DB, emails []string) ([]models.Teacher, error) {
	var teacherList []models.Teacher
	for _, email := range emails {
		var teacher models.Teacher
		err := db.Model(models.Teacher{}).
			Preload("Students").
			Where("email = ?", email).
			First(&teacher).
			Error
		if err != nil {
			return nil, errors.New("Teacher not found")
		}
		teacherList = append(teacherList, teacher)
	}

	return teacherList, nil
}

func RegisterStudentsToTeacher(db *gorm.DB, teacher *models.Teacher, students []models.Student) error {
	// TODO: Prevent adding repeated entries of students
	teacher.Students = append(teacher.Students, students...)
	return db.Updates(teacher).Error
}
