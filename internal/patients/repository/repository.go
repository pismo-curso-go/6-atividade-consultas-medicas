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

func (r *PatientRepository) FindByEmail(ctx context.Context, email string) (*patients.PatientDomain, error) {
	query := `SELECT id, name, email, password FROM patients WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email)

	var id int
	var name, emailDB, password string
	if err := row.Scan(&id, &name, &emailDB, &password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, ErrFailedQueryExec
	}
	return patients.NewPatientDomainFromDB(id, name, emailDB, password), nil
}
