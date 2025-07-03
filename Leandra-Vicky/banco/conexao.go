package banco

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB *sqlx.DB
)

func CarregarEnv() error {
	err := godotenv.Load("../Leandra-Vicky/.env")
	if err != nil {
		return fmt.Errorf("erro ao carregar .env: %w", err)
	}
	return nil

}

func ConectarDB() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return fmt.Errorf("erro ao abrir conexão com o banco de dados: %w", err)
	}

	log.Println("Conexão com o banco de dados estabelecida")
	return nil
}

func FecharConexao() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

func AutoMigrate() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS appointments (
			id SERIAL PRIMARY KEY,
			patient_id VARCHAR(1000) NOT NULL,
			datetime TIMESTAMP NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS usuarios
			id SERIAL PRIMARY KEY,
			nome VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			senha VARCHAR(255) NOT NULL
		);`,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatalf("Erro ao criar tabelas: %v", err)
		}
	}

	log.Println("Tabelas criadas com sucesso")
}
