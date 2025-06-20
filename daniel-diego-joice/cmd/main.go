package main

import (
	"log"
	"saude-mais/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Arquivo .env não encontrado. Usando variáveis do ambiente.")
	}

	server := server.NewServer(server.InitDB())
	if err := server.Start(); err != nil {
		panic(err)
	}
}
