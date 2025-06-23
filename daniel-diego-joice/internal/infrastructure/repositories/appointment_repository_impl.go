package repositories

import (
	"context"
	"database/sql"
	"saude-mais/internal/domain/entities"
	"saude-mais/internal/domain/repositories"
)

type appointmentRepositoryImpl struct {
	db *sql.DB
}

func NewAppointmentRepository(db *sql.DB) repositories.AppointmentRepository {
	return &appointmentRepositoryImpl{db: db}
}

func (r *appointmentRepositoryImpl) Create(ctx context.Context, patient *entities.Appointment) error {

	return nil
}

func (r *appointmentRepositoryImpl) GetAll(ctx context.Context) ([]*entities.Appointment, error) {

	return nil, nil
}

func (r *appointmentRepositoryImpl) GetByPatientID(ctx context.Context, id int) (*entities.Appointment, error) {
	return nil, nil
}

func (r *appointmentRepositoryImpl) Delete(ctx context.Context, id int) error {
	return nil
}

func (r *appointmentRepositoryImpl) Update(ctx context.Context, appointment *entities.Appointment) error {
	return nil
}


