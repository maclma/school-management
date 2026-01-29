package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	Port       string
	AppEnv     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, using system env")
	}

	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		Port:       os.Getenv("SERVER_PORT"),
	}
	return config, nil
}
