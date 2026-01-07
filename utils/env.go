// utils/env.go
package utils

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func LoadEnv() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not found")
		}
	})
}

func GetEnv(key string) string {
	LoadEnv()
	return os.Getenv(key)
}
