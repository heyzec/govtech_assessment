package handlers

import (
	"testing"

	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/errors"
	"github.com/heyzec/govtech-assignment/internal/handlers"
	"github.com/heyzec/govtech-assignment/internal/helpers/config"
	"github.com/heyzec/govtech-assignment/internal/helpers/database"
	"github.com/heyzec/govtech-assignment/internal/params"
	"github.com/heyzec/govtech-assignment/seeds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestRegister(t *testing.T) {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	dataaccess.CreateStudent(db, seeds.StudentJon)
	dataaccess.CreateStudent(db, seeds.StudentHon)
	dataaccess.CreateTeacher(db, seeds.TeacherKen)

	success(t, db)
	studentNotFound(t, db)
}

func success(t *testing.T, db *gorm.DB) {
	params := &params.RegisterStudentsParams{
		TeacherEmail:  "teacherken@gmail.com",
		StudentEmails: []string{"studentjon@gmail.com", "studenthon@gmail.com"},
	}

	res, err := handlers.HandleRegisterStudents(params, db)
	assert.Nil(t, res)
	assert.Nil(t, err)
}

func studentNotFound(t *testing.T, db *gorm.DB) {
	params := &params.RegisterStudentsParams{
		TeacherEmail:  "teacherken@gmail.com",
		StudentEmails: []string{"nosuchstudent@gmail.com", "studenthon@gmail.com"},
	}

	res, err := handlers.HandleRegisterStudents(params, db)
	assert.Nil(t, res)
	var customErr *errors.NotFoundError
	require.ErrorAs(t, err, &customErr)
}
