package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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
	// UNIDEAL WORKAROUND: Unable to use go test without it changing the directory
	// https://stackoverflow.com/questions/23847003/golang-tests-and-working-directory
	cwd, _ := os.Getwd()
	if strings.HasSuffix(cwd, "/tests/handlers") {
		os.Chdir("../..")
		cwd, _ = os.Getwd()
	}

	env := os.Getenv("GO_ENV")
	var configFileName string

	if env == "" {
		// This running tests within VSCode possible
		env = "testing"
		os.Setenv("GO_ENV", "testing")
	}

	switch env {
	case "development":
		configFileName = ".env.development"
	case "testing":
		configFileName = ".env.testing"
	default:
		log.Fatal("Please ensure GO_ENV is either 'development' or 'testing'")
	}

	cwd, _ = filepath.EvalSymlinks(cwd)
	configFilePath := filepath.Join(cwd, configFileName)
	if _, err := os.Stat(configFilePath); err != nil {
		log.Fatalf("Config file does not exist: %s\n", configFileName)
	}
	err := godotenv.Load(configFilePath)
	if err != nil {
		log.Panic(err.Error())
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
