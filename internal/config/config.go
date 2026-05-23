package config

import (
	"fmt"
	"os"
)

// Config holds all application configuration.
type Config struct {
	ServerPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	RedisHost string
	RedisPort string

	JWTSecret string
	JWTExpiry string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	cfg := &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "turf_user"),
		DBPassword: getEnv("DB_PASSWORD", "turf_password"),
		DBName:     getEnv("DB_NAME", "turf_booking"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),

		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnv("REDIS_PORT", "6379"),

		JWTSecret: getEnv("JWT_SECRET", ""),
		JWTExpiry: getEnv("JWT_EXPIRY", "24h"),
	}

	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET environment variable is required")
	}

	return cfg, nil
}

// DBConnectionString returns the PostgreSQL DSN.
func (c *Config) DBConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
