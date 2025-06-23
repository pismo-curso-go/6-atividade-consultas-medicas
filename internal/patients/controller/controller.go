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
}

func NewPatientController(useCase *usecase.PatientUseCase) *PatientController {
	return &PatientController{
		useCases: useCase,
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

	return config.ResponseMessageJSON(c, http.StatusCreated, "patient registered successfully")
}
