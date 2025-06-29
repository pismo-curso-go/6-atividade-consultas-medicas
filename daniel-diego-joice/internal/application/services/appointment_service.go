package services

import (
	"context"
	"saude-mais/internal/domain/entities"
	"saude-mais/internal/domain/repositories"
	"strconv"
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
	appointments, err := s.AppointmentRepo.GetAllByPatientID(ctx, patientID)
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
	appointment := &entities.Appointment{
		PatientID: patientID,
		DateTime:  dateTime,
	}
	return s.AppointmentRepo.Create(appointment)
}

func (s *AppointmentService) GetAllByPatientID(patientID int) ([]*entities.Appointment, error) {
	ctx := context.Background()
	return s.AppointmentRepo.GetAllByPatientID(ctx, patientID)
}

func (s *AppointmentService) GetByID(ctx context.Context, id string) (*entities.Appointment, error) {
	return s.AppointmentRepo.GetByID(ctx, id)
}

func (s *AppointmentService) UpdateAppointment(idStr string, newDate time.Time) error {
	ctx := context.Background()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	appointment := &entities.Appointment{
		ID:       id,
		DateTime: newDate,
	}

	return s.AppointmentRepo.Update(ctx, appointment)
}

func (s *AppointmentService) DeleteAppointment(ctx context.Context, id int) error {
	return s.AppointmentRepo.Delete(ctx, id)
}
