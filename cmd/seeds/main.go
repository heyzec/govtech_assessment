package main

import (
	"github.com/heyzec/govtech-assignment/internal/helpers/config"
	"github.com/heyzec/govtech-assignment/internal/helpers/database"
	"github.com/heyzec/govtech-assignment/seeds"
)

func main() {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)
	seeds.RunSeed(db)
}
