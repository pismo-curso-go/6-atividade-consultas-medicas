package repositories

import (
	"context"
	"database/sql"
	"log"
	"saude-mais/internal/domain/entities"
	"saude-mais/internal/domain/repositories"
)

type appointmentRepositoryImpl struct {
	db *sql.DB
}

func NewAppointmentRepository(db *sql.DB) repositories.AppointmentRepository {
	return &appointmentRepositoryImpl{db: db}
}

func (r *appointmentRepositoryImpl) Create(ctx context.Context, appointment *entities.Appointment) error {
	query := `
		INSERT INTO appointment (patient_id, date_time)
		VALUES ($1, $2)
		RETURNING id
		`
	err := r.db.QueryRowContext(
		ctx,
		query,
		appointment.PatientID,
		appointment.DateTime,
	).Scan(&appointment.ID)

	if err != nil {
		log.Printf("Erro ao criar consulta: %v", err)
		return err
	}

	return nil
}

func (r *appointmentRepositoryImpl) GetAll(ctx context.Context) ([]*entities.Appointment, error) {

	return nil, nil
}

func (r *appointmentRepositoryImpl) GetByPatientID(ctx context.Context, id int) ([]*entities.Appointment, error) {
	query := `SELECT id, patient_id, date_time FROM appointment WHERE patient_id = $1`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []*entities.Appointment

	for rows.Next() {
		var a entities.Appointment
		err := rows.Scan(&a.ID, &a.PatientID, &a.DateTime)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, &a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return appointments, nil
}

func (r *appointmentRepositoryImpl) Delete(ctx context.Context, id int) error {
	return nil
}

func (r *appointmentRepositoryImpl) Update(ctx context.Context, appointment *entities.Appointment) error {
	return nil
}
