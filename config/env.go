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
		port = "3000"
	}

	dbVariables := initDatabaseVariables()

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret"
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

func initDatabaseVariables() *DatabaseVariables {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "health_clinic_db"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	return &DatabaseVariables{
		host:     dbHost,
		user:     dbUser,
		password: dbPassword,
		port:     dbPort,
		name:     dbName,
	}
}
