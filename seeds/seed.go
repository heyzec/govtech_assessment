package seeds

import (
	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/models"
	"gorm.io/gorm"
)

func RunSeed(db *gorm.DB) {

	// var AllTeachers = []models.Teacher{
	// 	*TeacherKen,
	// 	*TeacherJoe,
	// }
	TeacherKen, _ = dataaccess.CreateTeacher(db, TeacherKen)
	TeacherJoe, _ = dataaccess.CreateTeacher(db, TeacherJoe)

	// Set-up relations
	StudentCommon1.Teachers = []models.Teacher{*TeacherKen, *TeacherJoe}
	StudentCommon2.Teachers = []models.Teacher{*TeacherKen, *TeacherJoe}
	StudentOnlyUnderTeacherKen.Teachers = []models.Teacher{*TeacherKen}

	var AllStudents = []models.Student{
		*StudentJon,
		*StudentHon,
		*StudentCommon1,
		*StudentCommon2,
		*StudentOnlyUnderTeacherKen,
		*StudentMary,
	}

	for _, student := range AllStudents {
		dataaccess.CreateStudent(db, &student)
	}
}
