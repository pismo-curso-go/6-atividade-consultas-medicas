package repository

import (
	"context"
	"database/sql"
	"errors"
	"healthclinic/internal/patients"
)

type PatientRepository struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) *PatientRepository {
	return &PatientRepository{
		db: db,
	}
}

func (r *PatientRepository) Save(ctx context.Context, data *patients.PatientDomain) error {
	query := `
		INSERT INTO patients (name, email, password)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.ExecContext(ctx, query, data.Name(), data.Email(), data.Password())
	if err != nil {
		return errors.New("database error while trying to save new patient")
	}

	return nil
}
