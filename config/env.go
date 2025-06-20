package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	errEnvFileNotFoundMessage = "Arquivo .env não encontrado. Usando variaveis com valores padrões"
)

type EnvVariables struct {
	port string
}

func InitEnvVariables() *EnvVariables {
	if err := godotenv.Load(); err != nil {
		log.Println(errEnvFileNotFoundMessage)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return &EnvVariables{
		port: port,
	}
}

func (ev *EnvVariables) Port() string {
	return ev.port
}
