package config

import (
	"fmt"
	"gin-quickstart/utils"
	"strconv"
)

type JWTConfig struct {
	SecretKey string
	AccessTokenExpireMinutes int
	RefreshTokenExpireDays int
}

// load jwt config
func LoadJWTConfig(envLoader func(string) string) (*JWTConfig, error) {
	accessMinutes, _ := strconv.Atoi(utils.GetValueOrDefault(envLoader("JWT_ACCESS_TOKEN_EXPIRE_MINUTES"), "15"))
	refreshDays, _ := strconv.Atoi(utils.GetValueOrDefault(envLoader("JWT_REFRESH_TOKEN_EXPIRE_DAYS"), "7"))

	config := &JWTConfig{
		SecretKey:                envLoader("JWT_SECRET"),
		AccessTokenExpireMinutes: accessMinutes,
		RefreshTokenExpireDays:   refreshDays,
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