package services

import (
	"context"
	"saude-mais/internal/application/dto"
	"saude-mais/internal/domain/repositories"
	"saude-mais/internal/utils"
)

type AuthService struct {
	patientRepo repositories.PatientRepository
	jwtSecret   string
}

func NewAuthService(patientRepo repositories.PatientRepository, jwtSecret string) *AuthService {
	return &AuthService{
		patientRepo: patientRepo,
		jwtSecret:   jwtSecret,
	}
}

func (s *AuthService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	patient, err := s.patientRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if patient == nil || !utils.CheckPasswordHash(req.Password, patient.Password) {
		return nil, utils.ErrInvalidCredentials
	}

	token, err := utils.GenerateJWT(patient.ID, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: token}, nil
}
