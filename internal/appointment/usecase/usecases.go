package usecase

import (
	"context"
	"healthclinic/internal/appointment"
	"healthclinic/internal/appointment/dto"
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

func (a *AppointmentUseCase) ListByPatientID(ctx context.Context, patientID int) ([]dto.AppointmentResponse, error) {
	appointments, err := a.repository.ListByPatientID(ctx, patientID)
	if err != nil {
		return nil, err
	}

	appointmentList := []dto.AppointmentResponse{}
	for _, aEntity := range appointments {
		res := dto.AppointmentResponse{
			ID:        aEntity.ID,
			PatientID: aEntity.PatientID,
			Date:      aEntity.Date,
		}
		appointmentList = append(appointmentList, res)
	}
	return appointmentList, nil
}

func (a *AppointmentUseCase) Create(ctx context.Context, patientID int, date time.Time) error {
	if date.Before(time.Now()) {
		return appointment.ErrAppointmentInvalidDate
	}

	conflict, err := a.repository.ExistsByPatientIDAndDate(ctx, patientID, date)
	if err != nil {
		return err
	}
	if conflict {
		return appointment.ErrAppointmentDoubleBooking
	}

	appointment, err := appointment.NewAppointmentDomain(0, date, patientID)
	if err != nil {
		return err
	}

	if err := a.repository.Save(ctx, appointment); err != nil {
		return err
	}

	return nil
}

func (a *AppointmentUseCase) Cancel(ctx context.Context, patientID, appointmentID int) error {
	found, err := a.repository.FindByIDAndPatientID(ctx, appointmentID, patientID)
	if err != nil {
		return err
	}
	if !found {
		return appointment.ErrAppointmentNotFoundOrUnauthorized
	}
	return a.repository.Delete(ctx, appointmentID)
}
