package repositories

import (
	"context"
	"saude-mais/internal/domain/entities"
)

type AppointmentRepository interface {
	Create(ctx context.Context, patient *entities.Appointment) error
	GetAll(ctx context.Context) ([]*entities.Appointment, error)
	GetByPatientID(ctx context.Context, id int) (*entities.Appointment, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, appointment *entities.Appointment) error
}