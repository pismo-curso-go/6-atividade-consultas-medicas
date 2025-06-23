package controller

import (
	"healthclinic/config"
	"healthclinic/internal/appointment/dto"
	"healthclinic/internal/appointment/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppointmentController struct {
	useCases *usecase.AppointmentUseCase
}

func NewAppointmentController(useCases *usecase.AppointmentUseCase) *AppointmentController {
	return &AppointmentController{
		useCases: useCases,
	}
}

func (p *AppointmentController) Create(c echo.Context) error {
	var payload dto.CreateAppointmentRequest
	if err := c.Bind(&payload); err != nil {
		return err
	}

	if err := p.useCases.Create(
		c.Request().Context(),
		payload.PatientID,
		payload.Date,
	); err != nil {
		return config.ResponseMessageJSON(c, http.StatusBadRequest, err.Error())
	}

	return config.ResponseMessageJSON(c, http.StatusCreated, "appointment created")
}

func (p *AppointmentController) ListByPatientID(c echo.Context) error {
	var payload dto.ListAppointmentRequest
	if err := c.Bind(&payload); err != nil {
		return err
	}

	appointments, err := p.useCases.ListByPatientID(
		c.Request().Context(),
		payload.PatientID,
	)
	if err != nil {
		return config.ResponseMessageJSON(c, http.StatusBadRequest, err.Error())
	}

	return config.ResponseJSON(c, http.StatusOK, appointments)
}
