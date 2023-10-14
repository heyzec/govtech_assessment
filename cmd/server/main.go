package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/heyzec/govtech-assignment/internal/config"
	"github.com/heyzec/govtech-assignment/internal/database"
	"github.com/heyzec/govtech-assignment/internal/handlers"
)

func main() {
    r := chi.NewRouter()
	r.Use(middleware.Logger)

    cfg, _ := config.LoadEnv()
    db := database.GetGormDB(cfg)

    r.Route("/api", func(r chi.Router) {
        r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
            handlers.RegisterStudents(w, r, db)
        })
        r.Get("/commonstudents", func(w http.ResponseWriter, r *http.Request) {
            handlers.HandleCommonStudents(w, r, db)
        })
    })

	http.ListenAndServe(fmt.Sprintf(":%d", cfg.ServerPort), r)
}

