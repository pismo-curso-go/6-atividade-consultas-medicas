package handlers

import (
	"saude-mais/internal/application/dto"
	"saude-mais/internal/application/services"
	"saude-mais/internal/utils"

	"github.com/labstack/echo/v4"
)

type AppointmentHandler struct {
	patientService *services.PatientService
}

func NewAppointmentHandler(patientService *services.PatientService, authService *services.AuthService) *PatientHandler {
	return &PatientHandler{
		patientService: patientService,
	}
}

func (h *AppointmentHandler) Create(c echo.Context) error {
	var req dto.CreateAppointmentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(utils.ErrValidationFailed.Code, utils.ErrValidationFailed)
	}

	if req.DateTime.IsZero() {
		return c.JSON(utils.ErrInvalidDateTime.Code, utils.ErrInvalidDateTime)	
	}

	// TO DO

	return nil
}