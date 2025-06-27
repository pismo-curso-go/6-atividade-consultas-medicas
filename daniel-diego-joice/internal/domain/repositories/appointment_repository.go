package repositories

import (
	"context"
	"saude-mais/internal/domain/entities"
)

type AppointmentRepository interface {
	Create(appointment *entities.Appointment) error
	GetAllByPatientID(ctx context.Context, id int) ([]*entities.Appointment, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, appointment *entities.Appointment) error
}
