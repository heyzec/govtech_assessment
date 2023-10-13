package seeds

import (
    "gorm.io/gorm"

    "github.com/heyzec/govtech-assignment/internal/dataaccess"
    "github.com/heyzec/govtech-assignment/internal/models"
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
    }

    for _, student := range students {
        dataaccess.CreateStudent(db, &student)
    }
}

