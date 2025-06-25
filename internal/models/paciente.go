package models

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Paciente struct {
	ID           int
	Nome         string
	Email        string
	PasswordHash string
}

func CriarPaciente(db *sql.DB, nome, email, senha string) error {
	var id int
	err := db.QueryRow("SELECT id FROM pacientes WHERE email = $1", email).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
    return err
}
	if id != 0 {
		return errors.New("Paciente já cadastrado")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		"INSERT INTO pacientes (name, email, password_hash) VALUES ($1, $2, $3)",
		nome, email, string(hash),
	)
	return err
}
