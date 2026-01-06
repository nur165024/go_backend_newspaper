package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type RedisConfig struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

func loadRedisConfig() *RedisConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		fmt.Println("Redis Host is required")
		os.Exit(1)
	}
	
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379" // default Redis port
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisPassword == "" {
		redisPassword = "" // default Redis password
	}


	redisDB := os.Getenv("REDIS_DB")
	if redisDB == "" {
		redisDB = "0" // default Redis DB
	}
	redisDBInt, _ := strconv.Atoi(redisDB)


	return &RedisConfig{
		RedisHost:     redisHost,
		RedisPort:     redisPort,
		RedisPassword: redisPassword,
		RedisDB:       redisDBInt,
	}
	
}

func GetRedisConfig() *RedisConfig {
	return loadRedisConfig()
}