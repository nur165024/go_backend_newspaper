package config

import (
	"fmt"
	"gin-quickstart/utils"
)

type JWTConfig struct {
	SecretKey string
	ExpiresIn string
}

// load jwt config
func LoadJWTConfig(envLoader func(string) string) (*JWTConfig, error) {
	config := &JWTConfig{
		SecretKey: envLoader("JWT_SECRET_KEY"),
		ExpiresIn: utils.GetValueOrDefault(envLoader("JWT_EXPIRES_IN"), "24h"),
	}
	err := validateJWTConfig(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func GetJWTConfig() (*JWTConfig, error) {
	return LoadJWTConfig(utils.GetEnv)
}

// validation jwt config
func validateJWTConfig(config *JWTConfig) error {
	if config.SecretKey == "" {
		return fmt.Errorf("JWT secret key is required")
	}
	return nil
}