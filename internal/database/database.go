package database

import (
	"fmt"
	"log"

	"github.com/heyzec/govtech-assignment/internal/config"
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
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	log.Printf("Connected to database: %s\n", cfg.DBName)
	return db
}

//func getSqlDB() *sql.DB{
//    dsn := buildDSN()
//    db, err := sql.Open("postgres", dsn)
//	if err != nil {
//        log.Fatalf("Failed to connect to default database: %v\n", err)
//	}
//    return db
//}
