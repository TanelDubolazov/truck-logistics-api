package config

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the project's root directory
func LoadEnv() {
	rootPath, err := filepath.Abs("./") // Adjust path as needed
	if err != nil {
		log.Fatalf("Error getting root path: %v", err)
	}

	envPath := filepath.Join(rootPath, ".env")

	err = godotenv.Load(envPath)
	if err != nil {
		log.Println("No .env file found, using system environment variables.")
	}
	log.Println("✅ Environment variables loaded.")
}
