package main

import (
	"api/banco"
	"api/router"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := banco.CarregarEnv(); err != nil {
		log.Fatal(err)
	}

	if err := banco.ConectarDB(); err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	banco.AutoMigrate()
	defer banco.FecharConexao()

	router.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(server.ListenAndServe())
}
