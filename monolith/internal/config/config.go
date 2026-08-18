package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// DSN = Data Source Name
	DBDSN string
}

func LoadConfig() *Config {
	// load file .env if there is any
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables, err: %w", err)
	}

	// gowallet_user:123456@tcp(localhost:3306)/gowallet?parseTime=true
	dsn := os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME") + "?parseTime=true"

	return &Config{
		DBDSN: dsn,
	}
}
