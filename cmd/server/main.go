package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/heyzec/govtech-assignment/internal/api"
	"github.com/heyzec/govtech-assignment/internal/config"
	"github.com/heyzec/govtech-assignment/internal/database"
	"github.com/heyzec/govtech-assignment/internal/handlers"
	"github.com/heyzec/govtech-assignment/internal/params"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	r.Route("/api", func(r chi.Router) {
		r.Post("/register", api.WrapHandler(db, handlers.HandleRegisterStudents, params.RegisterStudentsParseFrom))
		r.Get("/commonstudents", api.WrapHandler(db, handlers.HandleCommonStudents, params.CommonStudentsParseFrom))
		r.Post("/suspend", api.WrapHandler(db, handlers.HandleSuspend, params.SuspendParamsParseFrom))
	})

	http.ListenAndServe(fmt.Sprintf(":%d", cfg.ServerPort), r)
}
