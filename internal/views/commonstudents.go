package views

import (
	"github.com/heyzec/govtech-assignment/internal/models"
)

type CommonStudentsView struct {
	StudentEmails []string `json:"students"`
}

func CommonStudentsViewFrom(students []models.Student) CommonStudentsView {
	emails := []string{}
	for _, student := range students {
		emails = append(emails, student.Email)
	}

	view := CommonStudentsView{}
	view.StudentEmails = emails
	return view
}
