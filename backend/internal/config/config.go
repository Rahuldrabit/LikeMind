package config

import (
	"os"
	"strconv"
)

type Config struct {
	Environment   string
	Port          string
	DatabaseURL   string
	RedisURL      string
	JWTSecret     string
	OpenAIAPIKey  string
	VectorDBURL   string
	CORSOrigins   []string
}

func Load() *Config {
	return &Config{
		Environment:  getEnv("ENVIRONMENT", "development"),
		Port:         getEnv("PORT", "8080"),
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://postgres:password@localhost/likemind?sslmode=disable"),
		RedisURL:     getEnv("REDIS_URL", "redis://localhost:6379"),
		JWTSecret:    getEnv("JWT_SECRET", "your-super-secret-jwt-key-change-in-production"),
		OpenAIAPIKey: getEnv("OPENAI_API_KEY", ""),
		VectorDBURL:  getEnv("VECTOR_DB_URL", "http://localhost:6333"),
		CORSOrigins:  []string{getEnv("CORS_ORIGINS", "http://localhost:3000")},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
