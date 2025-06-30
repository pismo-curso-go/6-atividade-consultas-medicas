package postgres

import (
	"PacientesEAgendamento/internal/model"
	"PacientesEAgendamento/internal/repository"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresPatientRepository struct {
	db *pgxpool.Pool
}

// NewPostgresPatientRepository cria uma nova instância do repositório de pacientes.
func NewPostgresPatientRepository(db *pgxpool.Pool) repository.PatientRepository {
	return &postgresPatientRepository{db: db}
}

func (r *postgresPatientRepository) Create(ctx context.Context, patient *model.Patient) error {
	query := `INSERT INTO patients (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(ctx, query, patient.Name, patient.Email, patient.PasswordHash).Scan(&patient.ID, &patient.CreatedAt, &patient.UpdatedAt)

	if err != nil {
		var pgErr *pgconn.PgError
		// Verifica se o erro é de violação de constraint única (código 23505)
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return errors.New("email já cadastrado")
		}
		return err
	}
	return nil
}

func (r *postgresPatientRepository) FindByEmail(ctx context.Context, email string) (*model.Patient, error) {
	query := `SELECT id, name, email, password_hash, created_at, updated_at FROM patients WHERE email = $1`
	patient := &model.Patient{}
	err := r.db.QueryRow(ctx, query, email).Scan(&patient.ID, &patient.Name, &patient.Email, &patient.PasswordHash, &patient.CreatedAt, &patient.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("paciente não encontrado")
		}
		return nil, err
	}
	return patient, nil
}
