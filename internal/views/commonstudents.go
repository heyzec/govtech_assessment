package views

import (
	"github.com/heyzec/govtech-assignment/internal/models"
)

type commonStudentsView struct {
    StudentEmails []string `json:"students"`
}

func CommonStudentsViewFrom(students []models.Student) commonStudentsView {
    emails := []string{}
    for _, student := range students {
        emails = append(emails, student.Email)
    }

    view := commonStudentsView{}
    view.StudentEmails = emails
    return view
}

