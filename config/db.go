package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	errDatabaseNotConnectedMessage      = "Não foi possível se conectar ao banco de dados"
	errDatabaseInvalidConnectionMessage = "Erro na validação da conexão com o banco de dados"
)

func InitDB(env *EnvVariables) *sql.DB {
	strConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.db.host,
		env.db.user,
		env.db.password,
		env.db.name,
		env.db.port)

	database, err := sql.Open("postgres", strConn)

	if err != nil {
		log.Fatal(errDatabaseNotConnectedMessage+":", err)
	}

	if err = database.Ping(); err != nil {
		log.Fatal(errDatabaseInvalidConnectionMessage+":", err)
	}

	return database
}
