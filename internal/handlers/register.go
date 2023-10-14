package handlers

import (
	"log"
	"net/http"

	"github.com/heyzec/govtech-assignment/internal/api"
	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/json"
	"gorm.io/gorm"
)

type registerStudentsParams struct {
	TeacherEmail  string   `json:"teacher"`
	StudentEmails []string `json:"students"`
}

func HandleRegisterStudents(r *http.Request, db *gorm.DB) (*api.Response, error) {

	// Parse request body
	var registerStudentsParams registerStudentsParams
	err := json.DecodeParams(r.Body, &registerStudentsParams)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	students, err := dataaccess.FindStudentsByEmail(db, registerStudentsParams.StudentEmails)
	if err != nil {
		return nil, err
	}

	teachers, err := dataaccess.FindTeachersByEmail(db, []string{registerStudentsParams.TeacherEmail})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dataaccess.RegisterStudentsToTeacher(db, &teachers[0], students)

	return nil, nil
}
