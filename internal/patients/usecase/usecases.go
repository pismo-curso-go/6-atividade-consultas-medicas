package usecase

import (
	"context"
	"healthclinic/config"
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
	if password != "" {
		hashedPw, err := config.HashStr(password)
		if err != nil {
			return patients.ErrPatientInvalidPassword
		}
		password = hashedPw
	}

	existing, err := p.repository.FindByEmail(ctx, email)
	if err != nil {
		return err
	}
	if existing != nil {
		return patients.ErrPatientAlreadyExists
	}

	newPatient := patients.NewPatientDomain(name, email, password)
	if err := newPatient.Validate(); err != nil {
		return err
	}

	if err := p.repository.Save(ctx, newPatient); err != nil {
		return err
	}

	return nil
}

func (p *PatientUseCase) LoginPatientUseCase(ctx context.Context, email, password string) (*patients.PatientDomain, error) {
	patient, err := p.repository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, patients.ErrPatientInvalidPassword
	}
	if err := config.Compare(patient.Password(), password); err != nil {
		return nil, patients.ErrPatientInvalidPassword
	}
	return patient, nil
}
