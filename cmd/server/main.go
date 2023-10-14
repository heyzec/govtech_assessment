package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/heyzec/govtech-assignment/internal/helpers/config"
	"github.com/heyzec/govtech-assignment/internal/helpers/database"
	"github.com/heyzec/govtech-assignment/internal/routes"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	routes.SetupRoutes(r, db)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.ServerPort), r)
}
