package postgres

import (
	"PacientesEAgendamento/internal/repository"
	"context"
	"time"

	"PacientesEAgendamento/internal/model"
	"PacientesEAgendamento/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresAppointmentRepository struct {
	db *pgxpool.Pool
}

func NewPostgresAppointmentRepository(db *pgxpool.Pool) repository.AppointmentRepository {
	return &postgresAppointmentRepository{db: db}
}

func (r *postgresAppointmentRepository) Create(ctx context.Context, app *model.Appointment) error {
	query := `
		INSERT INTO appointments (patient_id, appointment_date, status)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query, app.PatientID, app.AppointmentDate, app.Status).
		Scan(&app.ID, &app.CreatedAt, &app.UpdatedAt)
	return err
}

func (r *postgresAppointmentRepository) ExistsByPatientAndDateTime(ctx context.Context, patientID string, dateTime time.Time) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM appointments WHERE patient_id = $1 AND appointment_date = $2 AND status = 'SCHEDULED')`
	err := r.db.QueryRow(ctx, query, patientID, dateTime).Scan(&exists)
	return exists, err
}

func (r *postgresAppointmentRepository) FindByPatientID(ctx context.Context, patientID string) ([]model.Appointment, error) {
	var appointments []model.Appointment
	query := `SELECT id, patient_id, appointment_date, status, created_at, updated_at FROM appointments WHERE patient_id = $1 ORDER BY appointment_date DESC`
	rows, err := r.db.Query(ctx, query, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var app model.Appointment
		if err := rows.Scan(&app.ID, &app.PatientID, &app.AppointmentDate, &app.Status, &app.CreatedAt, &app.UpdatedAt); err != nil {
			return nil, err
		}
		appointments = append(appointments, app)
	}

	return appointments, nil
}

func (r *postgresAppointmentRepository) FindByID(ctx context.Context, appointmentID string) (*model.Appointment, error) {
	app := &model.Appointment{}
	query := `SELECT id, patient_id, appointment_date, status, created_at, updated_at FROM appointments WHERE id = $1`
	err := r.db.QueryRow(ctx, query, appointmentID).Scan(&app.ID, &app.PatientID, &app.AppointmentDate, &app.Status, &app.CreatedAt, &app.UpdatedAt)
	if err == pgx.ErrNoRows {
		return nil, nil // Retorna nil, nil se não encontrar, para ser tratado no usecase.
	}
	return app, err
}

func (r *postgresAppointmentRepository) UpdateStatus(ctx context.Context, appointmentID string, status model.AppointmentStatus) error {
	query := `UPDATE appointments SET status = $1, updated_at = NOW() WHERE id = $2`
	cmdTag, err := r.db.Exec(ctx, query, status, appointmentID)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return pgx.ErrNoRows // Usa um erro padrão para indicar que nenhuma linha foi afetada.
	}
	return nil
}
