package usecase

import (
	"context"
	"healthclinic/internal/appointment"
	"healthclinic/internal/appointment/repository"
	"time"
)

type AppointmentUseCase struct {
	repository *repository.AppointmentRepository
}

func NewAppointmentUseCase(repo *repository.AppointmentRepository) *AppointmentUseCase {
	return &AppointmentUseCase{
		repository: repo,
	}
}

func (a *AppointmentUseCase) ListByPatientID(ctx context.Context, patientID int) ([]appointment.AppointmentDomain, error) {
	appointments, err := a.repository.ListByPatientID(ctx, patientID)
	if err != nil {
		return nil, err
	}

	appointmentList := []appointment.AppointmentDomain{}
	for _, aEntity := range appointments {
		aDomain, _ := appointment.NewAppointmentDomain(aEntity.ID, aEntity.Date, aEntity.PatientID)
		appointmentList = append(appointmentList, *aDomain)
	}
	return appointmentList, nil
}

func (a *AppointmentUseCase) Create(ctx context.Context, patientID int, date time.Time) error {
	appointment, err := appointment.NewAppointmentDomain(0, date, patientID)
	if err != nil {
		return err
	}

	if err := a.repository.Save(ctx, appointment); err != nil {
		return err
	}

	return nil
}
