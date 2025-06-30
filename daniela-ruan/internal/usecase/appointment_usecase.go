package usecase

import (
	"context"
	"errors"
	"time"

	"PacientesEAgendamento/internal/model"
	"PacientesEAgendamento/internal/repository"
)

type AppointmentUsecase struct {
	appointmentRepo repository.AppointmentRepository
}

func NewAppointmentUsecase(repo repository.AppointmentRepository) *AppointmentUsecase {
	return &AppointmentUsecase{appointmentRepo: repo}
}

func (uc *AppointmentUsecase) CreateAppointment(ctx context.Context, patientID string, appointmentDate time.Time) (*model.Appointment, error) {
	if appointmentDate.Before(time.Now()) {
		return nil, errors.New("consulta no passado não é permitida")
	}

	exists, err := uc.appointmentRepo.ExistsByPatientAndDateTime(ctx, patientID, appointmentDate)
	if err != nil {
		return nil, errors.New("erro ao verificar agendamentos existentes")
	}
	if exists {
		return nil, errors.New("paciente já possui uma consulta neste horário")
	}

	appointment := &model.Appointment{
		PatientID:       patientID,
		AppointmentDate: appointmentDate,
		Status:          model.StatusScheduled,
	}

	if err := uc.appointmentRepo.Create(ctx, appointment); err != nil {
		return nil, errors.New("erro ao salvar agendamento")
	}

	return appointment, nil
}

func (uc *AppointmentUsecase) ListByPatient(ctx context.Context, patientID string) ([]model.Appointment, error) {
	return uc.appointmentRepo.FindByPatientID(ctx, patientID)
}

func (uc *AppointmentUsecase) CancelAppointment(ctx context.Context, patientIDFromToken, appointmentID string) error {
	appointment, err := uc.appointmentRepo.FindByID(ctx, appointmentID)
	if err != nil {
		return errors.New("erro ao buscar consulta")
	}
	if appointment == nil || appointment.PatientID != patientIDFromToken {
		return errors.New("consulta não encontrada ou acesso não autorizado")
	}
	return uc.appointmentRepo.UpdateStatus(ctx, appointmentID, model.StatusCanceled)
}
