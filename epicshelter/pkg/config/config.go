package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Password string
	Db       int
	TTL      int
}

func LoadConfig() (*Config, error) {
	db, err := strconv.Atoi(getEnvOrDefault("CACHE_DATABASE", "0"))
	if err != nil {
		return nil, fmt.Errorf("invalid database: %v", err)
	}

	ttl, err := strconv.Atoi(getEnvOrDefault("CACHE_TTL", "300"))
	if err != nil {
		return nil, fmt.Errorf("invalid ttl value: %v", err)
	}

	return &Config{
		Host:     getEnvOrDefault("CACHE_URL", "localhost:6379"),
		Password: getEnvOrDefault("CACHE_PASSWORD", ""),
		Db:       db,
		TTL:      ttl,
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
