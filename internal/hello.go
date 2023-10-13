package main

import (
	"net/http"
    "fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/heyzec/govtech-assignment/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func registerStudents(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("welcome again"))
    return
}

func main() {
    r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

    r.Route("/api", func(r chi.Router) {
        r.Post("/register", registerStudents)
    })

    dsn := "host=localhost user=postgres password=postgres dbname=govtech_assignment port=5432 sslmode=disable TimeZone=Asia/Singapore"
    db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    var student []models.Student
    _ = db.Model(models.Student{}).Find(&student)


    fmt.Println(student)


	http.ListenAndServe(":3000", r)
}
