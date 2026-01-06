package cmd

import (
	"fmt"
	"gin-quickstart/config"
	"gin-quickstart/pkg/database"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"
)

func Migrate() {
	fmt.Println("Running database migrations...")

	dbCnf := config.GetDatabaseConfig()

	// database connection
	dbConnection := &database.DatabaseConfig{
		Host:     dbCnf.DBHost,
		Port:     dbCnf.DBPort,
		User:     dbCnf.DBUser,
		Password: dbCnf.DBPass,
		DbName:   dbCnf.DBName,
	}
	db := database.NewDatabaseConnection(dbConnection)
	defer db.Close()
	
	// migration direction
	migrationDirs := []string{
		"migration/users",
		"migration/categories",
	}

	for _, dir := range migrationDirs {
        fmt.Printf("Running migrations from: %s\n", dir)
        
        files, err := filepath.Glob(filepath.Join(dir, "*.sql"))
        if err != nil {
            log.Printf("Error reading migration files from %s: %v", dir, err)
            continue
        }
        
        for _, file := range files {
					migrationName := strings.Replace(filepath.Base(file), ".sql", "", 1)
            
            // Check if migration already executed
            if isMigrationExecuted(db, migrationName) {
                fmt.Printf("⏭️  Skipping (already executed): %s\n", file)
                continue
            }
						
            fmt.Printf("Executing: %s\n", file)
            
            content, err := ioutil.ReadFile(file)
            if err != nil {
                log.Printf("Error reading file %s: %v", file, err)
                continue
            }
            
            _, err = db.Exec(string(content))
            if err != nil {
                log.Printf("Error executing migration %s: %v", file, err)
                continue
            }
            
            fmt.Printf("✅ Successfully executed: %s\n", file)
        }
    }
    
    fmt.Println("🎉 All migrations completed!")
}

func createMigrationsTable(db *sqlx.DB) {
    query := `
    CREATE TABLE IF NOT EXISTS migrations (
        id SERIAL PRIMARY KEY,
        migration_name VARCHAR(255) UNIQUE NOT NULL,
        executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`
    
    db.Exec(query)
}

func isMigrationExecuted(db *sqlx.DB, migrationName string) bool {
    var count int
    query := "SELECT COUNT(*) FROM migrations WHERE migration_name = $1"
    db.Get(&count, query, migrationName)
    return count > 0
}

func markMigrationExecuted(db *sqlx.DB, migrationName string) {
    query := "INSERT INTO migrations (migration_name) VALUES ($1)"
    db.Exec(query, migrationName)
}