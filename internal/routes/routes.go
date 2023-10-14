package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/heyzec/govtech-assignment/internal/api"
	"github.com/heyzec/govtech-assignment/internal/handlers"
	"github.com/heyzec/govtech-assignment/internal/params"
	"gorm.io/gorm"
)

func SetupRoutes(r chi.Router, db *gorm.DB) {
	r.Route("/api", func(r chi.Router) {
		r.Post("/register", api.WrapHandler(db, handlers.HandleRegisterStudents, params.RegisterStudentsParseFrom))
		r.Get("/commonstudents", api.WrapHandler(db, handlers.HandleCommonStudents, params.CommonStudentsParseFrom))
		r.Post("/suspend", api.WrapHandler(db, handlers.HandleSuspend, params.SuspendParamsParseFrom))
	})
}
