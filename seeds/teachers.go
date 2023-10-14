package seeds

import (
    "gorm.io/gorm"

    "github.com/heyzec/govtech-assignment/internal/dataaccess"
    "github.com/heyzec/govtech-assignment/internal/models"
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


