package seeds

import (
	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/models"
	"gorm.io/gorm"
)

func RunSeed(db *gorm.DB) {

	// Set-up relations
	StudentCommon1.Teachers = []models.Teacher{*TeacherKen, *TeacherJoe}
	StudentCommon2.Teachers = []models.Teacher{*TeacherKen, *TeacherJoe}
	StudentOnlyUnderTeacherKen.Teachers = []models.Teacher{*TeacherKen}

	var AllTeachers = []models.Teacher{
		*TeacherKen,
		*TeacherJoe,
	}

	var AllStudents = []models.Student{
		*StudentJon,
		*StudentHon,
		*StudentCommon1,
		*StudentCommon2,
		*StudentOnlyUnderTeacherKen,
	}

	for _, teacher := range AllTeachers {
		dataaccess.CreateTeacher(db, &teacher)
	}
	for _, student := range AllStudents {
		dataaccess.CreateStudent(db, &student)
	}
}
