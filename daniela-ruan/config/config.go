package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL  string
	JWTSecretKey string
}

// LoadConfig carrega as configurações do arquivo .env ou das variáveis de ambiente.
func LoadConfig() *Config {
	// Carrega o arquivo .env, ignorando o erro se ele não existir (útil para produção)
	godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL não foi definida")
	}

	jwtKey := os.Getenv("JWT_SECRET_KEY")
	if jwtKey == "" {
		log.Fatal("JWT_SECRET_KEY não foi definida")
	}

	return &Config{
		DatabaseURL:  dbURL,
		JWTSecretKey: jwtKey,
	}
}
