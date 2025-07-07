package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret string
}

type ServerConfig struct {
	Port string
}

func Load() (*Config, error) {
	// Load .env file if it exists
	godotenv.Load()

	dbHost := "localhost"
	if os.Getenv("DOCKER_ENV") == "true" {
		dbHost = "postgres"
	}

	config := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", dbHost),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "saude_mais_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "your-super-secret-jwt-key"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8000"),
		},
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
