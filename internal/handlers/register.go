package handlers

import (
	"log"

	"github.com/heyzec/govtech-assignment/internal/api"
	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/params"
	"gorm.io/gorm"
)

func HandleRegisterStudents(params *params.RegisterStudentsParams, db *gorm.DB) (*api.Response, error) {
	students, err := dataaccess.FindStudentsByEmail(db, params.StudentEmails)
	if err != nil {
		return nil, err
	}

	teachers, err := dataaccess.FindTeachersByEmail(db, []string{params.TeacherEmail})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dataaccess.RegisterStudentsToTeacher(db, &teachers[0], students)

	return nil, nil
}
