package services

import "saude-mais/internal/domain/repositories"

type AppointmentService struct {
	AppointmentRepo repositories.AppointmentRepository
}

func NewAppointmentService(appointmentRepo repositories.AppointmentRepository) *AppointmentService {
	return &AppointmentService{
		AppointmentRepo: appointmentRepo,
	}
}
