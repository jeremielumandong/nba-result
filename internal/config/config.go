// Package config handles application configuration management.
package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds all configuration values for the application.
type Config struct {
	ServerAddress string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	IdleTimeout   time.Duration
	NBAAPIURL     string
}

// Load creates and returns a new Config instance with values from environment variables or defaults.
func Load() (*Config, error) {
	cfg := &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		ReadTimeout:   getDurationEnv("READ_TIMEOUT", 15*time.Second),
		WriteTimeout:  getDurationEnv("WRITE_TIMEOUT", 15*time.Second),
		IdleTimeout:   getDurationEnv("IDLE_TIMEOUT", 60*time.Second),
		NBAAPIURL:     getEnv("NBA_API_URL", "https://api.nba.com/v1"),
	}

	return cfg, nil
}

// getEnv retrieves an environment variable or returns a default value if not set.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getDurationEnv retrieves a duration environment variable or returns a default value.
func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	if seconds, err := strconv.Atoi(value); err == nil {
		return time.Duration(seconds) * time.Second
	}

	if duration, err := time.ParseDuration(value); err == nil {
		return duration
	}

	return defaultValue
}