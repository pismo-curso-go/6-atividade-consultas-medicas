package services

import (
	"context"
	"saude-mais/internal/application/dto"
	"saude-mais/internal/domain/entities"
	"saude-mais/internal/domain/repositories"
	"saude-mais/internal/utils"
)

type PatientService struct {
	patientRepo repositories.PatientRepository
}

func NewPatientService(patientRepo repositories.PatientRepository) *PatientService {
	return &PatientService{
		patientRepo: patientRepo,
	}
}

func (s *PatientService) Register(ctx context.Context, req *dto.RegisterPatientRequest) error {
	exists, err := s.patientRepo.EmailExists(ctx, req.Email)
	if err != nil {
		return err
	}
	if exists {
		return utils.ErrPatientAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	patient := &entities.Patient{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}
	return s.patientRepo.Create(ctx, patient)
}

func (s *PatientService) Delete(id string) error {
	return s.patientRepo.Delete(id)
}
