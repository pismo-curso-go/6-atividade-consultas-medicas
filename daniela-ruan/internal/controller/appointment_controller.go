package controller

import (
	"net/http"
	"time"

	"PacientesEAgendamento/internal/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AppointmentController struct {
	appointmentUsecase *usecase.AppointmentUsecase
}

func NewAppointmentController(uc *usecase.AppointmentUsecase) *AppointmentController {
	return &AppointmentController{appointmentUsecase: uc}
}

type CreateAppointmentRequest struct {
	DateTime string `json:"datetime"`
}

func (ac *AppointmentController) Create(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	patientID := claims["sub"].(string)

	var req CreateAppointmentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Message: "Corpo da requisição inválido", Code: http.StatusBadRequest})
	}

	appointmentDate, err := time.Parse("2006-01-02T15:04:05", req.DateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Message: "Formato de data inválido, use AAAA-MM-DDTHH:MM:SS", Code: http.StatusBadRequest})
	}

	appointment, err := ac.appointmentUsecase.CreateAppointment(c.Request().Context(), patientID, appointmentDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Message: err.Error(), Code: http.StatusBadRequest})
	}

	return c.JSON(http.StatusCreated, appointment)
}

func (ac *AppointmentController) List(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	patientID := claims["sub"].(string)

	appointments, err := ac.appointmentUsecase.ListByPatient(c.Request().Context(), patientID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, APIError{Message: "Erro ao listar consultas", Code: http.StatusInternalServerError})
	}

	return c.JSON(http.StatusOK, appointments)
}

func (ac *AppointmentController) Delete(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	patientIDFromToken := claims["sub"].(string)
	appointmentID := c.Param("id")

	err := ac.appointmentUsecase.CancelAppointment(c.Request().Context(), patientIDFromToken, appointmentID)
	if err != nil {
		if err.Error() == "consulta não encontrada ou acesso não autorizado" {
			return c.JSON(http.StatusForbidden, APIError{Message: err.Error(), Code: http.StatusForbidden})
		}
		return c.JSON(http.StatusInternalServerError, APIError{Message: "Erro ao cancelar consulta", Code: http.StatusInternalServerError})
	}

	return c.NoContent(http.StatusNoContent)
}
