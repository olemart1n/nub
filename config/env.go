// Package config returns environment variables
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	StorageZone string
	Hostname    string
	UploadKey   string
	PullZone    string
	DatabaseURL string
}

func LoadEnvConfig() EnvConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envs := EnvConfig{
		StorageZone: os.Getenv("BUNNY_STORAGE_ZONE"),
		Hostname:    os.Getenv("BUNNY_STORAGE_HOST"),
		UploadKey:   os.Getenv("BUNNY_UPLOAD_KEY"),
		PullZone:    os.Getenv("BUNNY_PULL_ZONE"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	return envs
}
