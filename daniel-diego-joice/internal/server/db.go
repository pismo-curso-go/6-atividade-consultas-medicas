package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
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
		log.Fatal("Database connection failed:", err)
	}

	if err = database.Ping(); err != nil {
		log.Fatal("Failed to verify database connection:", err)
	}

	return database
}
