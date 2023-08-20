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
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		return nil, fmt.Errorf("MONGO_URI environment variable is not set")
	}

	// Add more configuration loading logic here

	config := &AppConfig{
		MongoURI: mongoURI,
		// Initialize other configuration fields here
	}

	return config, nil
}
