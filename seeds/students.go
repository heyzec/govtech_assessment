package seeds

import (
	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/models"
	"gorm.io/gorm"
)

func CreateStudents(db *gorm.DB) {
	var students = []models.Student{
		{
			Email: "guy@gmail.com",
		},
		{
			Email: "studentjon@gmail.com",
		},
		{
			Email: "studenthon@gmail.com",
		},
		{
			Email: "commonstudent1@gmail.com",
		},
		{
			Email: "commonstudent2@gmail.com",
		},
		{
			Email: "student_only_under_teacher_ken@gmail.com",
		},
	}

	for _, student := range students {
		dataaccess.CreateStudent(db, &student)
	}
}
