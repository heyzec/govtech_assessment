package handlers

import (
	"fmt"
	"log"
	"regexp"

	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/helpers/api"
	"github.com/heyzec/govtech-assignment/internal/helpers/emailutils"
	"github.com/heyzec/govtech-assignment/internal/models"
	"github.com/heyzec/govtech-assignment/internal/params"
	"github.com/heyzec/govtech-assignment/internal/views"
	"gorm.io/gorm"
)

func HandleRetrieveForNotifications(params *params.RetrieveForNotificationParams, db *gorm.DB) (*api.Response, error) {
	teachers, err := dataaccess.FindTeachersByEmail(db, []string{params.TeacherEmail})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	teacher := teachers[0]
	registeredStudents := teacher.Students

	mentionedRegex := regexp.MustCompile(fmt.Sprintf("@%s", emailutils.EmailRegex))

	mentionedEmails := []string{}
	for _, mentioned := range mentionedRegex.FindAllString(params.Notification, -1) {
		email := mentioned[1:]
		mentionedEmails = append(mentionedEmails, email)
	}

	mentionedStudents, err := dataaccess.FindStudentsByEmail(db, mentionedEmails)
	if err != nil {
		return nil, err
	}

	seenStudentIDs := make(map[uint]bool)
	outputStudents := []models.Student{}
	for _, student := range registeredStudents {
		_, value := seenStudentIDs[student.ID]
		if !value && !student.Suspended {
			outputStudents = append(outputStudents, student)
			seenStudentIDs[student.ID] = true
		}
	}
	for _, student := range mentionedStudents {
		_, value := seenStudentIDs[student.ID]
		if !value && !student.Suspended {
			outputStudents = append(outputStudents, student)
			seenStudentIDs[student.ID] = true
		}
	}

	view := views.RetrieveForNotificationsViewFrom(outputStudents)
	return &api.Response{
		JSONPayload: &view,
	}, nil
}
