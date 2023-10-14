package handlers

import (
	"testing"

	"github.com/heyzec/govtech-assignment/internal/errors"
	"github.com/heyzec/govtech-assignment/internal/handlers"
	"github.com/heyzec/govtech-assignment/internal/helpers/config"
	"github.com/heyzec/govtech-assignment/internal/helpers/database"
	"github.com/heyzec/govtech-assignment/internal/params"
	"github.com/stretchr/testify/require"
)

func TestRegisterSuccess(t *testing.T) {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	params := &params.RegisterStudentsParams{
		TeacherEmail:  "teacherken@gmail.com",
		StudentEmails: []string{"studentjon@gmail.com", "studenthon@gmail.com"},
	}

	res, err := handlers.HandleRegisterStudents(params, db)
	require.Nil(t, res)
	require.Nil(t, err)
}

func TestRegisterStudentNotFound(t *testing.T) {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	params := &params.RegisterStudentsParams{
		TeacherEmail:  "teacherken@gmail.com",
		StudentEmails: []string{"nosuchstudent@gmail.com", "studenthon@gmail.com"},
	}

	res, err := handlers.HandleRegisterStudents(params, db)
	require.Nil(t, res)
	var customErr *errors.NotFoundError
	require.ErrorAs(t, err, &customErr)
}
