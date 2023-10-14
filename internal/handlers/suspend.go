package handlers

import (
	"log"

	"github.com/heyzec/govtech-assignment/internal/api"
	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/params"
	"gorm.io/gorm"
)

func HandleSuspend(params *params.SuspendParams, db *gorm.DB) (*api.Response, error) {
	students, err := dataaccess.FindStudentsByEmail(db, []string{params.StudentEmail})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	student := students[0]

	dataaccess.SuspendStudent(db, &student)
	return nil, nil
}
