// Package config returns environment variables
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	StorageRegion           string
	StorageZone             string
	StoragePassword         string
	StoragePasswordReadOnly string
	PullZone                string
	DatabaseURL             string
	RedisURL                string
	PORT                    string
}

func LoadEnvConfig() EnvConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envs := EnvConfig{
		StorageRegion:           os.Getenv("BUNNY_STORAGE_REGION"),
		StorageZone:             os.Getenv("BUNNY_STORAGE_ZONE"),
		StoragePassword:         os.Getenv("BUNNY_STORAGE_PASSWORD"),
		StoragePasswordReadOnly: os.Getenv("BUNNY_STORAGE_PASSWWORD_READ_ONLY"),
		PullZone:                os.Getenv("BUNNY_PULL_ZONE"),
		DatabaseURL:             os.Getenv("DATABASE_URL"),
		RedisURL:                os.Getenv("REDIS_URL"),
		PORT:                    os.Getenv("PORT"),
	}

	return envs
}
