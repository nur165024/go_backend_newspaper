package main

import (
	"gin-quickstart/cmd"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		cmd.Migrate()
		return	
	}

	cmd.Server()
}

// main.go
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}
}
