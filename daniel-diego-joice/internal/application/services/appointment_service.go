package services

import (
	"context"
	"saude-mais/internal/domain/entities"
	"saude-mais/internal/domain/repositories"
	"time"
)

type AppointmentService struct {
	AppointmentRepo repositories.AppointmentRepository
}

func NewAppointmentService(appointmentRepo repositories.AppointmentRepository) *AppointmentService {
	return &AppointmentService{
		AppointmentRepo: appointmentRepo,
	}
}

func (s *AppointmentService) HasAppointmentAt(dateTime time.Time, patientID int) (bool, error) {
	ctx := context.Background()
	appointments, err := s.AppointmentRepo.GetByPatientID(ctx, patientID)
	if err != nil {
		return false, err
	}

	for _, a := range appointments {
		if a.DateTime.Equal(dateTime) {
			return true, nil
		}
	}
	return false, nil
}

func (s *AppointmentService) CreateAppointment(dateTime time.Time, patientID int) error {
	ctx := context.Background()

	appointment := &entities.Appointment{
		PatientID: patientID,
		DateTime:  dateTime,
	}

	return s.AppointmentRepo.Create(ctx, appointment)
}
