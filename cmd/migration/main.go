package main

import (
	"fmt"
	"log"

	"github.com/heyzec/govtech-assignment/internal/helpers/config"
	"github.com/heyzec/govtech-assignment/internal/helpers/database"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	cfg := config.LoadEnv()
	db := database.GetSQLDB(cfg)

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	steps, err := migrate.ExecMax(db, "postgres", migrations, migrate.Up, 0)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Applied %d migrations!\n", steps)
}
