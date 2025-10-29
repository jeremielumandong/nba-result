// Package config handles application configuration
package config

import (
	"os"
)

// Config holds all configuration for the application
type Config struct {
	Port   string
	NBAAPI string
}

// Load returns a Config struct with values from environment variables or defaults
func Load() *Config {
	return &Config{
		Port:   getEnvOrDefault("PORT", "8080"),
		NBAAPI: getEnvOrDefault("NBA_API_URL", "https://api.nba.com"),
	}
}

// getEnvOrDefault returns the value of an environment variable or a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}