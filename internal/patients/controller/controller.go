package controller

import (
	"healthclinic/config"
	"healthclinic/internal/patients/dto"
	"healthclinic/internal/patients/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PatientController struct {
	useCases *usecase.PatientUseCase
	env      *config.EnvVariables
}

func NewPatientController(useCase *usecase.PatientUseCase, env *config.EnvVariables) *PatientController {
	return &PatientController{
		useCases: useCase,
		env:      env,
	}
}

func (p *PatientController) RegisterPacient(c echo.Context) error {
	var payload dto.RegisterPatientRequest
	if err := c.Bind(&payload); err != nil {
		return err
	}

	if err := p.useCases.CreatePatientUseCase(
		c.Request().Context(),
		payload.Name,
		payload.Password,
		payload.Email,
	); err != nil {
		return config.ResponseMessageJSON(c, http.StatusBadRequest, err.Error())
	}

	return config.ResponseMessageJSON(c, http.StatusCreated, "Paciente cadastrado com sucesso")
}

func (p *PatientController) Login(c echo.Context) error {
	var payload dto.LoginPatientRequest
	if err := c.Bind(&payload); err != nil {
		return config.ResponseMessageJSON(c, http.StatusBadRequest, "Dados inválidos")
	}

	patient, err := p.useCases.LoginPatientUseCase(
		c.Request().Context(),
		payload.Email,
		payload.Password,
	)
	if err != nil {
		return config.ResponseMessageJSON(c, http.StatusUnauthorized, "Token inválido ou expirado")
	}

	secret := p.env.JwtSecret()
	token, err := config.GenerateJWT(patient.ID(), patient.Email(), secret)
	if err != nil {
		return config.ResponseMessageJSON(c, http.StatusInternalServerError, "Erro ao gerar token")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
