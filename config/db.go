package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const (
	errDatabaseNotConnectedMessage      = "Não foi possível se conectar ao banco de dados"
	errDatabaseInvalidConnectionMessage = "Erro na validação da conexão com o banco de dados"
)

func InitDB() *sql.DB {
	strConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	database, err := sql.Open("postgres", strConn)

	if err != nil {
		log.Fatal(errDatabaseNotConnectedMessage+":", err)
	}

	if err = database.Ping(); err != nil {
		log.Fatal(errDatabaseInvalidConnectionMessage+":", err)
	}

	return database
}
