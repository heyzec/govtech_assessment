package seeds

import (
	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/models"
	"gorm.io/gorm"
)

func CreatePeople(db *gorm.DB) {
	StudentJon := &models.Student{
		Email: "studentjon@gmail.com",
	}
	StudentHon := &models.Student{
		Email: "studenthon@gmail.com",
	}
	StudentCommon1 := &models.Student{
		Email: "commonstudent1@gmail.com",
	}
	StudentCommon2 := &models.Student{
		Email: "commonstudent2@gmail.com",
	}
	StudentOnlyUnderTeacherKen := &models.Student{
		Email: "student_only_under_teacher_ken@gmail.com",
	}

	var students = []models.Student{
		*StudentJon,
		*StudentHon,
		*StudentCommon1,
		*StudentCommon2,
		*StudentOnlyUnderTeacherKen,
	}

	TeacherKen := &models.Teacher{
		Email: "teacherken@gmail.com",
	}
	TeacherJoe := &models.Teacher{
		Email: "teacherjoe@gmail.com",
	}

	var teachers = []models.Teacher{
		*TeacherKen,
		*TeacherJoe,
	}

	// Set-up relations
	StudentCommon1.Teachers = []models.Teacher{*TeacherKen, *TeacherJoe}
	StudentCommon2.Teachers = []models.Teacher{*TeacherKen, *TeacherJoe}
	StudentOnlyUnderTeacherKen.Teachers = []models.Teacher{*TeacherKen}

	for _, teacher := range teachers {
		dataaccess.CreateTeacher(db, &teacher)
	}
	for _, student := range students {
		dataaccess.CreateStudent(db, &student)
	}
}
