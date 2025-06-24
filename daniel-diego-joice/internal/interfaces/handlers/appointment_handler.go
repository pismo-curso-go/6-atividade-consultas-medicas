package handlers

import (
	"net/http"
	"time"

	"saude-mais/internal/application/dto"
	"saude-mais/internal/application/services"
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

	user := c.Get("user").(dto.JWTClaims)
	patientID := user.UserID

	hasConflict, err := h.appointmentService.HasAppointmentAt(req.DateTime, patientID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrInternalServer)
	}
	if hasConflict {
		return c.JSON(utils.ErrAppointmentConflict.Code, utils.ErrAppointmentConflict)
	}

	err = h.appointmentService.CreateAppointment(req.DateTime, patientID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrInternalServer)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Consulta agendada com sucesso",
	})
}

func (h *AppointmentHandler) GetByID(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "get by id funcionando",
	})
}

func (h *AppointmentHandler) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "update funcionando",
	})
}

func (h *AppointmentHandler) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "delete funcionando",
	})
}
func (h *AppointmentHandler) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "get all funcionando",
	})
}
