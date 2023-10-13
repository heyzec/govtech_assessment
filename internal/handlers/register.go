package handlers

import (
	"log"
	"net/http"

    "gorm.io/gorm"

	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/json"
)

type registerStudentsParams struct {
    TeacherEmail  string `json:"teacher"`
    StudentEmails []string `json:"students"`
}


func RegisterStudents(w http.ResponseWriter, r *http.Request, db *gorm.DB) {

    // Parse request body
    var registerStudentsParams registerStudentsParams
    err := json.DecodeParams(r.Body, &registerStudentsParams)
    if err != nil {
        log.Println(err.Error())
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    students, _ := dataaccess.FindStudentsByEmail(db, registerStudentsParams.StudentEmails)
    if err != nil {
        log.Println(err)
        return
    }

    teachers, err := dataaccess.FindTeachersByEmail(db, []string{registerStudentsParams.TeacherEmail})
    if err != nil {
        log.Println(err)
        return
    }

    dataaccess.RegisterStudentsToTeacher(db, &teachers[0], students)

    w.WriteHeader(http.StatusNoContent)
    return
}

