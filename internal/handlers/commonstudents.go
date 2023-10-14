package handlers

import (
	"log"

	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/helpers/api"
	"github.com/heyzec/govtech-assignment/internal/helpers/arrayutils"
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

	idToStudents := make(map[uint]models.Student)
	var intersected []uint
	for i, teacher := range teachers {
		studentIds := []uint{}
		for _, student := range teacher.Students {
			studentIds = append(studentIds, student.ID)
			if _, value := idToStudents[student.ID]; !value {
				idToStudents[student.ID] = student
			}
		}
		if i == 0 {
			intersected = studentIds
		} else {
			intersected = arrayutils.Intersection(intersected, studentIds)
		}
	}

	studentsList := []models.Student{}
	for _, studentId := range intersected {
		studentsList = append(studentsList, idToStudents[studentId])
	}

	view := views.CommonStudentsViewFrom(studentsList)

	return &api.Response{
		JSONPayload: &view,
	}, nil
}
