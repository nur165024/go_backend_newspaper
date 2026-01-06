package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port       string
	ServerName string
	Mode       string
	Version    string
	Host       string
}

func loadServerConfig() *ServerConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// http port 
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Http Port is required")
		os.Exit(1)
	}

	// server name
	serverName := os.Getenv("SERVER_NAME")
	if serverName == "" {
		fmt.Println("Server Name is required")
		os.Exit(1)
	}

	// mode 
	mode := os.Getenv("MODE")
	if mode == "" {
		fmt.Println("Mode is required")
		os.Exit(1)
	}

	// version
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	// host
	host := os.Getenv("HOST")
	if host == "" {
		fmt.Println("Host is required")
		os.Exit(1)
	}

	return &ServerConfig{
		Port:       port,
		ServerName: serverName,
		Mode:       mode,
		Version:    version,
		Host:       host,
	}
}

func GetServerConfig() *ServerConfig {
	return loadServerConfig()
}