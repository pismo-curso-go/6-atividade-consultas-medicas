package controller

import (
	"healthclinic/config"
	"healthclinic/internal/appointment/dto"
	"healthclinic/internal/appointment/usecase"
	"net/http"
	"strconv"

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

func (a *AppointmentController) Create(c echo.Context) error {
	var payload dto.CreateAppointmentRequest
	if err := c.Bind(&payload); err != nil {
		return config.ResponseMessageJSON(c, http.StatusBadRequest, "Dados inválidos")
	}

	patientID, ok := c.Get("patient_id").(int)
	if !ok {
		return config.ResponseMessageJSON(c, http.StatusUnauthorized, "Token inválido ou expirado")
	}

	if err := a.useCases.Create(
		c.Request().Context(),
		patientID,
		payload.Date,
	); err != nil {
		return config.ResponseMessageJSON(c, http.StatusBadRequest, err.Error())
	}

	return config.ResponseMessageJSON(c, http.StatusCreated, "Consulta agendada com sucesso")
}

func (a *AppointmentController) ListByPatientID(c echo.Context) error {
	patientID, ok := c.Get("patient_id").(int)
	if !ok {
		return config.ResponseMessageJSON(c, http.StatusUnauthorized, "Token inválido ou expirado")
	}

	appointments, err := a.useCases.ListByPatientID(
		c.Request().Context(),
		patientID,
	)
	if err != nil {
		return config.ResponseMessageJSON(c, http.StatusBadRequest, err.Error())
	}

	return config.ResponseJSON(c, http.StatusOK, appointments)
}

func (a *AppointmentController) Cancel(c echo.Context) error {
	patientID, ok := c.Get("patient_id").(int)
	if !ok {
		return config.ResponseMessageJSON(c, http.StatusUnauthorized, "Token inválido ou expirado")
	}
	appointmentIDStr := c.Param("id")
	appointmentID, err := strconv.Atoi(appointmentIDStr)
	if err != nil {
		return config.ResponseMessageJSON(c, http.StatusBadRequest, "ID inválido")
	}

	err = a.useCases.Cancel(c.Request().Context(), patientID, appointmentID)
	if err != nil {
		return config.ResponseMessageJSON(c, http.StatusForbidden, err.Error())
	}

	return config.ResponseMessageJSON(c, http.StatusOK, "Consulta cancelada com sucesso")
}
