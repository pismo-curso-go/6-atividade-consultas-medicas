package main

import (
	"log"
	"os"
	"saudemais-api/internal/database"
	"saudemais-api/internal/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado. Variáveis de ambiente devem estar no sistema.")
	}
}

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco: %v", err)
	}
	defer db.Close()

	e := echo.New()

	e.POST("/register", handlers.Register(db))
    e.POST("/login", handlers.Login(db))
	// e.POST("/appointments", handlers.AgendarConsulta(db), middleware.Auth(db))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Servidor ouvindo na porta %s...", port)
	e.Logger.Fatal(e.Start(":" + port))

}
