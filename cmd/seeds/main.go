package main

import (
	"github.com/heyzec/govtech-assignment/internal/config"
	"github.com/heyzec/govtech-assignment/internal/database"
	"github.com/heyzec/govtech-assignment/seeds"
)

func main() {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)
	seeds.RunSeed(db)
}
