package main

import (
	"github.com/heyzec/govtech-assignment/seeds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=govtech_assignment port=5432 sslmode=disable TimeZone=Asia/Singapore"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	seeds.RunSeed(db)
}
