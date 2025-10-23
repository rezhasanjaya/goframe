package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string

	RedisHost string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using environment variables only")
	}

	cfg := &Config{
		AppPort: getEnv("APP_PORT", "8080"),

		DBHost: getEnv("DB_HOST", "127.0.0.1"),
		DBUser: getEnv("DB_USER", "root"),
		DBPass: getEnv("DB_PASS", ""),
		DBName: getEnv("DB_NAME", "goframe"),
		DBPort: getEnv("DB_PORT", "3306"),

		RedisHost: getEnv("REDIS_HOST", "127.0.0.1:6379"),
	}

	log.Println("✅ Config loaded")
	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
