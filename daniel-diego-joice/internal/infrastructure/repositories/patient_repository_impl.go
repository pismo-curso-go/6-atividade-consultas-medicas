package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"saude-mais/internal/domain/entities"
	"saude-mais/internal/domain/repositories"

	"time"
)

type patientRepositoryImpl struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) repositories.PatientRepository {
	return &patientRepositoryImpl{db: db}
}

func (r *patientRepositoryImpl) Create(ctx context.Context, patient *entities.Patient) error {
	query := `
        INSERT INTO patient (name, email, password, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5) 
        RETURNING id
    `

	now := time.Now()
	err := r.db.QueryRowContext(
		ctx,
		query,
		patient.Name,
		patient.Email,
		patient.Password,
		now,
		now,
	).Scan(&patient.ID)

	if err != nil {
		log.Printf("Erro ao criar paciente: %v", err)
		return err
	}

	patient.CreatedAt = now
	patient.UpdatedAt = now

	return nil
}

func (r *patientRepositoryImpl) GetByEmail(ctx context.Context, email string) (*entities.Patient, error) {
	query := `
		SELECT id, name, email, password, created_at, updated_at 
		FROM patient 
		WHERE email = $1
	`

	patient := &entities.Patient{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&patient.ID,
		&patient.Name,
		&patient.Email,
		&patient.Password,
		&patient.CreatedAt,
		&patient.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return patient, nil
}

func (r *patientRepositoryImpl) GetByID(ctx context.Context, id int) (*entities.Patient, error) {
	query := `
		SELECT id, name, email, password, created_at, updated_at 
		FROM patient 
		WHERE id = $1
	`

	patient := &entities.Patient{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&patient.ID,
		&patient.Name,
		&patient.Email,
		&patient.Password,
		&patient.CreatedAt,
		&patient.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return patient, nil
}

func (r *patientRepositoryImpl) EmailExists(ctx context.Context, email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM patient WHERE email = $1)`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *patientRepositoryImpl) Delete(id string) error {
	query := `DELETE FROM patients WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete patient: %w", err)
	}

	return nil
}
