package views

import (
	"github.com/heyzec/govtech-assignment/internal/models"
)

type RetrieveForNotificationsView struct {
	StudentEmails []string `json:"recipents"`
}

func RetrieveForNotificationsViewFrom(students []models.Student) RetrieveForNotificationsView {
	emails := []string{}
	for _, student := range students {
		emails = append(emails, student.Email)
	}

	view := RetrieveForNotificationsView{}
	view.StudentEmails = emails
	return view
}
