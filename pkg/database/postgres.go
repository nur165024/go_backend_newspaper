package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Host string
	Port string
	User string
	Password string
	DbName string
	SSLMode string
}

func NewDatabaseConnection(cfg *DatabaseConfig) *sqlx.DB  {
	// Default SSL mode
	if cfg.SSLMode == "" {
			cfg.SSLMode = "disable"
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SSLMode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Successfully connected to database")
	return db
}