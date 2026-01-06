package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPass string
}

func loadDatabaseConfig() *DatabaseConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// database host
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("Database Host is required!")
		os.Exit(1)
	}

	// database port
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("Database Port is required!")
		os.Exit(1)
	}
	
	// database name
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("Database Name is required!")
		os.Exit(1)
	}
	// database user
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("Database User is required!")
		os.Exit(1)
	}

	// database password
	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		fmt.Println("Database Password is required!")
	}

	return &DatabaseConfig{
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBName:     dbName,
		DBUser:     dbUser,
		DBPass: dbPass,
	}
}

func GetDatabaseConfig() *DatabaseConfig {
	return loadDatabaseConfig()
}