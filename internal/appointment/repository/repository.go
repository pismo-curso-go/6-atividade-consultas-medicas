package repository

import (
	"context"
	"database/sql"
	"healthclinic/internal/appointment"
	"time"

	"github.com/labstack/gommon/log"
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
		log.Error(err)
		return ErrFailedQueryExec
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
		log.Error(err)
		return nil, ErrFailedQueryExec
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
			return nil, ErrFailedRowScan
		}
		appointments = append(appointments, a)
	}

	if err := rows.Err(); err != nil {
		return nil, ErrInvalidIteration
	}

	return appointments, nil
}
