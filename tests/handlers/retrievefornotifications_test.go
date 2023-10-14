package handlers_test

import (
	"testing"

	"github.com/heyzec/govtech-assignment/internal/errors"
	"github.com/heyzec/govtech-assignment/internal/handlers"
	"github.com/heyzec/govtech-assignment/internal/helpers/config"
	"github.com/heyzec/govtech-assignment/internal/helpers/database"
	"github.com/heyzec/govtech-assignment/internal/params"
	"github.com/heyzec/govtech-assignment/internal/views"
	"github.com/stretchr/testify/require"
)

func TestRetrieveForNotificationsSuccess1(t *testing.T) {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	params := &params.RetrieveForNotificationParams{
		TeacherEmail: "teacherken@gmail.com",
		Notification: "Hello students! @studentjon@gmail.com @studenthon@gmail.com",
	}

	output, err := handlers.HandleRetrieveForNotifications(params, db)
	expected := &views.RetrieveForNotificationsView{
		StudentEmails: []string{
			"commonstudent1@gmail.com",
			"commonstudent2@gmail.com",
			"student_only_under_teacher_ken@gmail.com",
			"studentjon@gmail.com",
			"studenthon@gmail.com",
		},
	}

	require.Nil(t, err)
	require.Equal(t, output.JSONPayload, expected)
}

func TestRetrieveForNotificationsSuccess2(t *testing.T) {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	params := &params.RetrieveForNotificationParams{
		TeacherEmail: "teacherken@gmail.com",
		Notification: "Hey everybody",
	}

	output, err := handlers.HandleRetrieveForNotifications(params, db)
	expected := &views.RetrieveForNotificationsView{
		StudentEmails: []string{
			"commonstudent1@gmail.com",
			"commonstudent2@gmail.com",
			"student_only_under_teacher_ken@gmail.com",
		},
	}

	require.Nil(t, err)
	require.Equal(t, output.JSONPayload, expected)
}

func TestRetrieveForNotificationsTeacherDoesNotExist(t *testing.T) {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	params := &params.RetrieveForNotificationParams{
		TeacherEmail: "teacherdoesnotexist@gmail.com",
		Notification: "Hey everybody",
	}

	res, err := handlers.HandleRetrieveForNotifications(params, db)

	require.Nil(t, res)
	var customErr *errors.NotFoundError
	require.ErrorAs(t, err, &customErr)
}

func TestRetrieveForNotificationsRepeatedMentions(t *testing.T) {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	params := &params.RetrieveForNotificationParams{
		TeacherEmail: "teacherken@gmail.com",
		Notification: "@student_only_under_teacher_ken@gmail.com @student_only_under_teacher_ken@gmail.com @student_only_under_teacher_ken@gmail.com",
	}

	res, err := handlers.HandleRetrieveForNotifications(params, db)
	expected := &views.RetrieveForNotificationsView{
		StudentEmails: []string{
			"commonstudent1@gmail.com",
			"commonstudent2@gmail.com",
			"student_only_under_teacher_ken@gmail.com",
		},
	}

	require.Nil(t, err)
	require.Equal(t, res.JSONPayload, expected)
}
