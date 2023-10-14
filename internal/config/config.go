package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

// Refer to .env file
type Config struct {
	ServerPort int
	DBHostname string
	DBPort     int
	DBName     string
	DBUsername string
	DBPassword string
	DBSSLMode  string
}

func LoadEnv() *Config {
	env := os.Getenv("GO_ENV")
	var configFileName string
	switch env {
	case "development":
		configFileName = ".env.development"
	case "testing":
		configFileName = ".env.testing"
	default:
		log.Fatal("Please ensure GO_ENV is either 'development' or 'testing'")
	}

	cwd, _ := os.Getwd()
	cwd, _ = filepath.EvalSymlinks(cwd)
	err := godotenv.Load(filepath.Join(cwd, configFileName))
	if err != nil {
		log.Fatalf("Unable to load file: %s\n", configFileName)
	}

	var config Config
	config.ServerPort, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))
	config.DBHostname = os.Getenv("DB_HOSTNAME")
	config.DBPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	config.DBName = os.Getenv("DB_NAME")
	config.DBUsername = os.Getenv("DB_USERNAME")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBSSLMode = os.Getenv("DB_SSLMODE")

	return &config
}
