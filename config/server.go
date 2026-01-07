package config

import (
	"fmt"
	"gin-quickstart/utils"
)

type ServerConfig struct {
	Port       string
	ServerName string
	Mode       string
	Version    string
	Host       string
}

func LoadServerConfig(envLoader func(string) string) (*ServerConfig, error) {
	config := &ServerConfig{
		Port: utils.GetValueOrDefault(envLoader("PORT"), "8080"),
		ServerName: utils.GetValueOrDefault(envLoader("SERVER_NAME"), "Newspaper API"),
		Mode: utils.GetValueOrDefault(envLoader("MODE"), "development"),
		Version: utils.GetValueOrDefault(envLoader("VERSION"), "1.0.0"),
		Host: utils.GetValueOrDefault(envLoader("HOST"), "localhost"),
	}
	
	if err := validateServerConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func GetServerConfig() (*ServerConfig, error) {
	return LoadServerConfig(utils.GetEnv)
}

func validateServerConfig(config *ServerConfig) error {
	if config.Port == "" {
		return fmt.Errorf("port is required")
	}
	return nil
}