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

func HandleRegisterStudents(r *http.Request, db *gorm.DB) *api.Response {

	// Parse request body
	var registerStudentsParams registerStudentsParams
	err := json.DecodeParams(r.Body, &registerStudentsParams)
	if err != nil {
		log.Println(err.Error())
		return &api.Response{
			HTTPCode: http.StatusBadRequest,
		}
	}

	students, _ := dataaccess.FindStudentsByEmail(db, registerStudentsParams.StudentEmails)
	if err != nil {
		log.Println(err)
		return nil
	}

	teachers, err := dataaccess.FindTeachersByEmail(db, []string{registerStudentsParams.TeacherEmail})
	if err != nil {
		log.Println(err)
		return nil
	}

	dataaccess.RegisterStudentsToTeacher(db, &teachers[0], students)

	return &api.Response{
		HTTPCode: http.StatusNoContent,
	}
}
