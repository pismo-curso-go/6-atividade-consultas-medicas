package usecase

import (
	"PacientesEAgendamento/internal/model"
	"PacientesEAgendamento/internal/repository"
	"PacientesEAgendamento/internal/security"
	"context"
	"errors"
)

type PatientUsecase struct {
	patientRepo repository.PatientRepository
	jwtSecret   string
}

// NewPatientUsecase cria uma nova instância do use case de paciente.
func NewPatientUsecase(repo repository.PatientRepository, secret string) *PatientUsecase {
	return &PatientUsecase{
		patientRepo: repo,
		jwtSecret:   secret,
	}
}

// Register cria um novo paciente.
func (uc *PatientUsecase) Register(ctx context.Context, name, email, password string) (*model.Patient, error) {
	if name == "" || email == "" || password == "" {
		return nil, errors.New("nome, email e senha são obrigatórios")
	}

	hashedPassword, err := security.Hash(password)
	if err != nil {
		return nil, errors.New("erro ao processar senha")
	}

	patient := &model.Patient{
		Name:         name,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	if err := uc.patientRepo.Create(ctx, patient); err != nil {
		return nil, err
	}

	return patient, nil
}

// Login autentica um paciente e retorna um token JWT.
func (uc *PatientUsecase) Login(ctx context.Context, email, password string) (string, error) {
	patient, err := uc.patientRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("credenciais inválidas")
	}

	if err := security.VerifyPassword(patient.PasswordHash, password); err != nil {
		return "", errors.New("credenciais inválidas")
	}

	token, err := security.GenerateJWT(patient.ID, uc.jwtSecret)
	if err != nil {
		return "", errors.New("erro ao gerar token de acesso")
	}

	return token, nil
}
