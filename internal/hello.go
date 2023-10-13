package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/heyzec/govtech-assignment/internal/config"
	"github.com/heyzec/govtech-assignment/internal/dataaccess"
	"github.com/heyzec/govtech-assignment/internal/database"

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

    cfg, _ := config.LoadEnv()
    db := database.GetGormDB(cfg)

    students, _ := dataaccess.ListStudents(db)

    fmt.Println(students)


	http.ListenAndServe(fmt.Sprintf(":%d", cfg.ServerPort), r)
}
