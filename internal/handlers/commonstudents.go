package handlers

import (
	"log"
	"net/http"

	"gorm.io/gorm"

	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/json"
	"github.com/heyzec/govtech-assignment/internal/models"
)

type commonStudentsParams struct {
    StudentEmails []string `json:"students"`
}

type commonStudentsView struct {
    StudentEmails []string `json:"students"`
}

func ViewFrom(students []models.Student) commonStudentsView {
    emails := []string{}
    for _, student := range students {
        emails = append(emails, student.Email)
    }

    view := commonStudentsView{}
    view.StudentEmails = emails
    return view

}

func HandleCommonStudents(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Parse URL parameter
    teacherEmails := r.URL.Query()["teacher"]

    teachers, err := dataaccess.FindTeachersByEmail(db, teacherEmails)
    if err != nil {
        // TODO: Return 404
        log.Println(err)
        return
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

    view := ViewFrom(studentsList)

    raw, err := json.EncodeView(view)
    if err != nil {
        println("Error encoding view")
        return
    }


    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(raw)
    return
}

