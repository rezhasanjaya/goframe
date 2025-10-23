package config

import (
	"log"
	"os"
	"strconv"

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

	JWTSecret         string `json:"-"`
    AccessTokenTTLMin int    `json:"-"`
    RefreshTokenTTLH  int    `json:"-"`
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

		JWTSecret:         getEnv("JWT_SECRET", "replace_this_with_env_secret"),
        AccessTokenTTLMin: mustAtoi(getEnv("ACCESS_TOKEN_TTL_MIN", "15")),
        RefreshTokenTTLH:  mustAtoi(getEnv("REFRESH_TOKEN_TTL_H", "24")), // hours
	}

	log.Println("✅ Config loaded")
	return cfg
}

func mustAtoi(s string) int {
    v, _ := strconv.Atoi(s)
    if v == 0 {
        return 1
    }
    return v
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
