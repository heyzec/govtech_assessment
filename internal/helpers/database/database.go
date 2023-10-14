package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/heyzec/govtech-assignment/internal/helpers/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func buildDSN(cfg *config.Config) string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.DBHostname,
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode)
	return dsn
}

func GetGormDB(cfg *config.Config) *gorm.DB {
	dsn := buildDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	log.Printf("Connected to database: %s\n", cfg.DBName)
	return db
}

func GetSQLDB(cfg *config.Config) *sql.DB {
	dsn := buildDSN(cfg)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	return db
}
