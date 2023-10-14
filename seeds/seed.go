package seeds

import (
	"log"

	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/models"
	"gorm.io/gorm"
)

func RunSeed(db *gorm.DB) {

	// var AllTeachers = []models.Teacher{
	// 	*TeacherKen,
	// 	*TeacherJoe,
	// }
	TeacherKen, err := dataaccess.CreateTeacher(db, TeacherKen)
	if err != nil {
		log.Fatalln(err)
	}
	TeacherJoe, err := dataaccess.CreateTeacher(db, TeacherJoe)
	if err != nil {
		log.Fatalln(err)
	}

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
		_, err := dataaccess.CreateStudent(db, &student)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
