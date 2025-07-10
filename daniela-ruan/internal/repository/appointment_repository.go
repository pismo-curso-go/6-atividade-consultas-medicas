package repository

import (
	"PacientesEAgendamento/internal/model"
	"context"
	"time"
)

// AppointmentRepository define as operações de banco de dados para um agendamento.
type AppointmentRepository interface {
	Create(ctx context.Context, appointment *model.Appointment) error
	ExistsByPatientAndDateTime(ctx context.Context, patientID string, dateTime time.Time) (bool, error)
	FindByPatientID(ctx context.Context, patientID string) ([]model.Appointment, error)
	FindByID(ctx context.Context, appointmentID string) (*model.Appointment, error)
	UpdateStatus(ctx context.Context, appointmentID string, status model.AppointmentStatus) error
}
