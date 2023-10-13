package main

import (
    "gorm.io/gorm"
	"gorm.io/driver/postgres"

    "github.com/heyzec/govtech-assignment/seeds"
)

func main() {
    dsn := "host=localhost user=postgres password=postgres dbname=govtech_assignment port=5432 sslmode=disable TimeZone=Asia/Singapore"
    db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    seeds.CreateStudents(db)
}
