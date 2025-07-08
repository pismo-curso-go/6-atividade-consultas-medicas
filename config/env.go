package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	errEnvFileNotFoundMessage = "Arquivo .env não encontrado. Usando variaveis com valores padrões"
)

type DatabaseVariables struct {
	host     string
	user     string
	password string
	port     string
	name     string
}
type EnvVariables struct {
	db        *DatabaseVariables
	jwtSecret string
	port      string
}

func InitEnvVariables() *EnvVariables {
	if err := godotenv.Load(); err != nil {
		log.Println(errEnvFileNotFoundMessage)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is required")
	}

	dbVariables := initDatabaseVariables()

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	return &EnvVariables{
		port:      port,
		db:        dbVariables,
		jwtSecret: jwtSecret,
	}
}

func (ev *EnvVariables) Port() string {
	return ev.port
}

func (ev *EnvVariables) JwtSecret() string {
	return ev.jwtSecret
}

func initDatabaseVariables() *DatabaseVariables {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("DB_HOST environment variable is required")
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("DB_USER environment variable is required")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("DB_PASSWORD environment variable is required")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME environment variable is required")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatal("DB_PORT environment variable is required")
	}

	return &DatabaseVariables{
		host:     dbHost,
		user:     dbUser,
		password: dbPassword,
		port:     dbPort,
		name:     dbName,
	}
}
