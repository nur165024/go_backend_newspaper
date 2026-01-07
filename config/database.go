package config

import (
	"fmt"
	"gin-quickstart/utils"
	"log"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPass     string
}

func loadDatabaseConfig(envLoader func(string) string) (*DatabaseConfig, error) {
	config := &DatabaseConfig{
		DBHost: utils.GetValueOrDefault(envLoader("DB_HOST"), "localhost"),
		DBPort: utils.GetValueOrDefault(envLoader("DB_PORT"), "5432"),
		DBName: utils.GetValueOrDefault(envLoader("DB_NAME"), "go_newspaper"),
		DBUser: utils.GetValueOrDefault(envLoader("DB_USER"), "postgres"),
		DBPass: utils.GetValueOrDefault(envLoader("DB_PASS"), "1234"),
	}

	if err := validateDatabaseConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func GetDatabaseConfig() (*DatabaseConfig, error) {
	// Load .env file first
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}
	return loadDatabaseConfig(utils.GetEnv)
}

func validateDatabaseConfig(config *DatabaseConfig) error {
	if config.DBName == "" {
		return fmt.Errorf("database name is required")
	}
	if config.DBUser == "" {
		return fmt.Errorf("database user is required")
	}
	if config.DBPass == "" {
		return fmt.Errorf("database password is required")
	}
	return nil
}