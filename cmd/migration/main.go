package main

import (
	"fmt"
	"log"

	"github.com/heyzec/govtech-assignment/internal/config"
	"github.com/heyzec/govtech-assignment/internal/database"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	cfg := config.LoadEnv()
	fmt.Println(cfg)
	db := database.GetSQLDB(cfg)

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	steps, err := migrate.ExecMax(db, "postgres", migrations, migrate.Up, 0)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(steps)
}
