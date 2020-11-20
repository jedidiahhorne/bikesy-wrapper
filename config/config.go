package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.uber.org/config"
)

// Application ...
type Application struct {
	Name string
	Port string
}

// Profile provides routing for each safety preference ("standard", etc.) to correct OSRM instance
type Profile struct {
	Host string
}

// Profiles ...
type Profiles struct {
	Standard Profile
}

// Osrm includes all profiles
type Osrm struct {
	Profiles Profiles
}

// Redis stores connection info for elevation data
type Redis struct {
	URL string
}

// Configuration ...
type Configuration struct {
	Application Application
	Osrm Osrm
	Redis Redis
}

// LoadConfig reads development.yaml for now
func LoadConfig(logger *log.Logger) (*Configuration, error) {
	var c Configuration
	err := godotenv.Load()
	  if err != nil {
	    log.Fatal("Error loading .env file")
	  }

	path := os.Getenv("CONFIG")
	port := os.Getenv("PORT")
	cfg, err := config.NewYAML(config.File(path))
	if err != nil {
		return &c, err
	}

	if err := cfg.Get("").Populate(&c); err != nil {
		return &c, err
	}

	// use as secret only
	c.Redis.URL = os.Getenv("REDIS_URL")

	if port != "" {
		logger.Printf("Using custom port %v", port)
		c.Application.Port = port
	}
	return &c, nil
}
