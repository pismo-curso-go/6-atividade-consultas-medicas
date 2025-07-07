package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"saude-mais/internal/application/dto"
	"saude-mais/internal/application/services"
	"saude-mais/internal/domain/repositories"
	"saude-mais/internal/utils"

	"github.com/labstack/echo/v4"
)

type AppointmentHandler struct {
	appointmentService *services.AppointmentService
}

func NewAppointmentHandler(appointmentService *services.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		appointmentService: appointmentService,
	}
}

func (h *AppointmentHandler) Create(c echo.Context) error {
	var req dto.CreateAppointmentRequest
	var userID int = c.Get("user_id").(int)

	req.PatientID = userID

	if err := c.Bind(&req); err != nil {
		return c.JSON(utils.ErrValidationFailed.Code, utils.ErrValidationFailed)
	}

	if req.DateTime.IsZero() {
		return c.JSON(utils.ErrInvalidDateTime.Code, utils.ErrInvalidDateTime)
	}

	now := time.Now()
	if req.DateTime.Before(now) {
		return c.JSON(utils.ErrPastAppointment.Code, utils.ErrPastAppointment)
	}

	hasConflict, err := h.appointmentService.HasAppointmentAt(req.DateTime, req.PatientID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrInternalServer)
	}
	if hasConflict {
		return c.JSON(utils.ErrAppointmentConflict.Code, utils.ErrAppointmentConflict)
	}

	err = h.appointmentService.CreateAppointment(req.DateTime, req.PatientID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrInternalServer)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Consulta agendada com sucesso",
	})
}

func (h *AppointmentHandler) GetAllAppointmentByPatientID(c echo.Context) error {
	var userID int = c.Get("user_id").(int)

	appointments, err := h.appointmentService.GetAllByPatientID(userID)
	if err != nil {
		if apiErr, ok := err.(utils.APIError); ok {
			return c.JSON(apiErr.Code, apiErr)
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrInternalServer)
	}

	if len(appointments) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrAppointmentNotFound)
	}

	return c.JSON(http.StatusOK, appointments)
}

func (h *AppointmentHandler) Update(c echo.Context) error {
	id := c.Param("id")
	userID := c.Get("user_id").(int)

	var req dto.UpdateAppointmentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(utils.ErrValidationFailed.Code, utils.ErrValidationFailed)
	}

	if req.DateTime.IsZero() {
		return c.JSON(utils.ErrInvalidDateTime.Code, utils.ErrInvalidDateTime)
	}
	if req.DateTime.Before(time.Now()) {
		return c.JSON(utils.ErrPastAppointment.Code, utils.ErrPastAppointment)
	}

	ctx := c.Request().Context()
	appointment, err := h.appointmentService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrAppointmentNotFound)
	}
	if appointment.PatientID != userID {
		return c.JSON(http.StatusForbidden, utils.ErrForbidden)
	}

	hasConflict, err := h.appointmentService.HasAppointmentAt(req.DateTime, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrInternalServer)
	}
	if hasConflict {
		return c.JSON(utils.ErrAppointmentConflict.Code, utils.ErrAppointmentConflict)
	}

	err = h.appointmentService.UpdateAppointment(id, req.DateTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrInternalServer)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Consulta atualizada com sucesso",
	})
}

func (h *AppointmentHandler) GetByID(c echo.Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()
	appointment, err := h.appointmentService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrInternalServer)
	}
	if appointment == nil {
		return c.JSON(http.StatusNotFound, utils.ErrAppointmentNotFound)
	}

	return c.JSON(http.StatusOK, appointment)
}

type AppointmentService struct {
	AppointmentRepo repositories.AppointmentRepository
}

func (h *AppointmentHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	userID := c.Get("user_id").(int)

	ctx := c.Request().Context()
	appointment, err := h.appointmentService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrAppointmentNotFound)
	}
	if appointment.PatientID != userID {
		return c.JSON(http.StatusForbidden, utils.ErrForbidden)
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
	}

	err = h.appointmentService.DeleteAppointment(ctx, intID)
	if err != nil {
		fmt.Println("Tentando deletar appointment ID:", id)
		fmt.Println("UserID do token:", userID)
		fmt.Println("appointment.PatientID:", appointment.PatientID)
		return c.JSON(http.StatusInternalServerError, utils.ErrInternalServer)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Consulta cancelada com sucesso",
	})
}
