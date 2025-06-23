package repository

import (
	"context"
	"database/sql"
	"errors"
	"healthclinic/internal/appointment"
	"time"
)

type AppointmentEntity struct {
	ID        int
	PatientID int
	Date      time.Time
}

type AppointmentRepository struct {
	db *sql.DB
}

func NewAppointmentRepository(db *sql.DB) *AppointmentRepository {
	return &AppointmentRepository{
		db: db,
	}
}

func (r *AppointmentRepository) Save(ctx context.Context, data *appointment.AppointmentDomain) error {
	query := `
		INSERT INTO appointments (patient_id, date)
		VALUES ($1, $2)
	`

	_, err := r.db.ExecContext(ctx, query, data.PatientID(), data.Date())
	if err != nil {
		return errors.New("database error while trying to save new patient")
	}

	return nil
}

func (r *AppointmentRepository) ListByPatientID(ctx context.Context, patientID int) ([]AppointmentEntity, error) {
	query := `
		SELECT id, patient_id, date
		FROM appointments
		WHERE appointments.patient_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, patientID)
	if err != nil {
		return nil, errors.New("database err - 1")
	}
	defer rows.Close()

	appointments := []AppointmentEntity{}
	for rows.Next() {
		var a AppointmentEntity
		if err := rows.Scan(
			&a.ID,
			&a.PatientID,
			&a.Date,
		); err != nil {
			return nil, errors.New("database err - 2")
		}
		appointments = append(appointments, a)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New("database err - 3")
	}

	return appointments, nil
}
