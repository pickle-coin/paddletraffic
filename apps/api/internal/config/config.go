package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server     ServerConfig
	Database   DatabaseConfig
	Pagination PaginationConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type DatabaseConfig struct {
	URL            string
	MaxConnections int
	MigrationPath  string
}

type PaginationConfig struct {
	DefaultPageSize int
	MaxPageSize     int
}

func Load() (*Config, error) {
	dbURL := getEnv("DATABASE_URL", "postgresql://paddletraffic:password@localhost:5432/paddletraffic")

	return &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			ReadTimeout:  getDuration("READ_TIMEOUT", 15*time.Second),
			WriteTimeout: getDuration("WRITE_TIMEOUT", 15*time.Second),
			IdleTimeout:  getDuration("IDLE_TIMEOUT", 60*time.Second),
		},
		Database: DatabaseConfig{
			URL:            dbURL,
			MaxConnections: getEnvInt("DB_MAX_CONNS", 25),
			MigrationPath:  getEnv("MIGRATION_PATH", "file://internal/database/migrations"),
		},
		Pagination: PaginationConfig{
			DefaultPageSize: getEnvInt("DEFAULT_PAGE_SIZE", 50),
			MaxPageSize:     getEnvInt("MAX_PAGE_SIZE", 100),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
