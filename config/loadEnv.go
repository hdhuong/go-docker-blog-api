package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Warning: unable to load .env file: %v", err)
		log.Println("Falling back to system environment variables")
	}
}
