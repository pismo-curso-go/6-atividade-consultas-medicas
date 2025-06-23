package usecase

import (
	"context"
	"healthclinic/internal/patients"
	"healthclinic/internal/patients/repository"
)

type PatientUseCase struct {
	repository *repository.PatientRepository
}

func NewPatientUseCase(repo *repository.PatientRepository) *PatientUseCase {
	return &PatientUseCase{
		repository: repo,
	}
}

func (p *PatientUseCase) CreatePatientUseCase(ctx context.Context, name, password, email string) error {
	newPatient := patients.NewPatientDomain(name, email, password)
	if err := newPatient.Validate(); err != nil {
		return err
	}

	if err := p.repository.Save(ctx, newPatient); err != nil {
		return err
	}

	return nil
}
