package seeds

import (
	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/models"
	"gorm.io/gorm"
)

func CreateTeachers(db *gorm.DB) {
	var seedables = []models.Teacher{
		{
			Email: "teacherken@gmail.com",
		},
		{
			Email: "teacherjoe@gmail.com",
		},
	}

	for _, seedable := range seedables {
		dataaccess.CreateTeacher(db, &seedable)
	}
}
