package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


// Config has access to all the configuration variables needed for the application
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	frontendURL string

}


// LoadConfig function retrieves configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}
	// retrieve each variable from the environment
	config := &Config{
		DBHost:   os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		frontendURL: os.Getenv("FRONTEND_URL"),
	}

	// Check to see if any required variables are missing
	if config.DBPassword == "" {
		return nil, fmt.Errorf("DB_password is required in .env file")
	}
	if config.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required in .env file")
	}
	return config, nil
}

// getEnv replaces os.getenv with fallback values if env var is missing
func getEnv(key, defaultValue string) string  {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}