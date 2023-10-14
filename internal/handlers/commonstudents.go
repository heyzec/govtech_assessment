package handlers

import (
	"log"

	"github.com/heyzec/govtech-assignment/internal/api"
	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/models"
	"github.com/heyzec/govtech-assignment/internal/params"
	"github.com/heyzec/govtech-assignment/internal/views"
	"gorm.io/gorm"
)

func HandleCommonStudents(params *params.CommonStudentsParams, db *gorm.DB) (*api.Response, error) {

	teachers, err := dataaccess.FindTeachersByEmail(db, params.TeacherEmails)
	if err != nil {
		// TODO: Return 404
		log.Println(err)
		return nil, err
	}

	studentsMap := make(map[uint]*models.Student)
	for _, teacher := range teachers {
		for i, student := range teacher.Students {
			if i == 0 {
				studentsMap[student.ID] = &student
				continue
			}
			if _, value := studentsMap[student.ID]; !value {
				delete(studentsMap, student.ID)
			}
		}
	}

	studentsList := []models.Student{}
	for _, student := range studentsMap {
		studentsList = append(studentsList, *student)
	}

	view := views.CommonStudentsViewFrom(studentsList)

	return &api.Response{
		JSONPayload: &view,
	}, nil
}
