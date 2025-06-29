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

type AppointmentService struct {
	repo repositories.AppointmentRepository
}

func NewAppointmentRepository(db *sql.DB) repositories.AppointmentRepository {
	return &appointmentRepositoryImpl{db: db}
}

func (r *appointmentRepositoryImpl) Create(appointment *entities.Appointment) error {
	query := `
		INSERT INTO appointment (patient_id, date_time)
		VALUES ($1, $2)
		`
	_, err := r.db.Exec(query, appointment.PatientID, appointment.DateTime)
	if err != nil {
		log.Printf("Erro ao criar consulta: %v", err)
		return err
	}

	return nil
}

func (r *appointmentRepositoryImpl) GetAllByPatientID(ctx context.Context, id int) ([]*entities.Appointment, error) {
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
	query := `DELETE FROM appointment WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *appointmentRepositoryImpl) Update(ctx context.Context, appointment *entities.Appointment) error {
	query := `UPDATE appointment SET date_time = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, appointment.DateTime, appointment.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *appointmentRepositoryImpl) GetByID(ctx context.Context, id string) (*entities.Appointment, error) {
	query := `SELECT id, patient_id, date_time FROM appointment WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)

	var a entities.Appointment
	err := row.Scan(&a.ID, &a.PatientID, &a.DateTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Erro ao buscar consulta por ID: %v", err)
		return nil, err
	}

	return &a, nil
}
