package repository

import (
	"context"
	"database/sql"
	"healthclinic/internal/patients"

	"github.com/labstack/gommon/log"
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
		log.Error(err)
		return ErrFailedQueryExec
	}

	return nil
}
