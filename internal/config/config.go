package config

import (
	"fmt"
	"os"
)

type AppConfig struct {
	MongoURI string
	// Add more configuration fields as needed
}

func LoadConfig() (*AppConfig, error) {
	mongoURI := os.Getenv("MONGODB_URL")
	if mongoURI == "" {
		return nil, fmt.Errorf("MONGODB_URL environment variable is not set")
	}

	// Add more configuration loading logic here

	config := &AppConfig{
		MongoURI: mongoURI,
		// Initialize other configuration fields here
	}

	return config, nil
}
