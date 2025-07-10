package repository

import (
	"PacientesEAgendamento/internal/model"
	"context"
)

// PatientRepository define as operações de banco de dados para um paciente.
type PatientRepository interface {
	Create(ctx context.Context, patient *model.Patient) error
	FindByEmail(ctx context.Context, email string) (*model.Patient, error)
}
