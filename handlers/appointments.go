package handlers

import (
	"errors"
	"net/http"
	"saudemais-api/auth"
	"saudemais-api/database"
	"saudemais-api/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AppointmentRequest struct {
	DateTime string `json:"datetime"`
}

// getPatientIDFromToken é uma função auxiliar para extrair o ID do paciente
// das 'claims' do token JWT que o middleware de autenticação já validou e
// armazenou no contexto da requisição.
func getPatientIDFromToken(c echo.Context) (uint, error) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return 0, errors.New("token JWT inválido")
	}

	claims, ok := user.Claims.(*auth.JwtCustomClaims)
	if !ok {
		return 0, errors.New("claims JWT inválidas")
	}

	return claims.PatientID, nil
}

// CreateAppointment lida com a criação de uma nova consulta.
func CreateAppointment(c echo.Context) error {
	req := new(AppointmentRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error(), "code": http.StatusBadRequest})
	}

	// Validação 1: Converte e verifica o formato da data (RFC3339)
	appointmentTime, err := time.Parse(time.RFC3339, req.DateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Formato de data e hora inválido. Use AAAA-MM-DDTHH:MM:SSZ", "code": http.StatusBadRequest})
	}

	// Validação 2: Impede agendamentos em datas passadas.
	if appointmentTime.Before(time.Now()) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Consulta no passado não é permitida", "code": http.StatusBadRequest})
	}

	patientID, err := getPatientIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Acesso não autorizado", "code": http.StatusUnauthorized})
	}

	// Validação 3: Impede que o mesmo paciente agende duas consultas no mesmo horário.
	var existingAppointment models.Appointment
	if result := database.DB.Where("patient_id = ? AND date_time = ?", patientID, appointmentTime).First(&existingAppointment); result.Error == nil {
		return c.JSON(http.StatusConflict, map[string]interface{}{"message": "Já existe uma consulta agendada neste horário", "code": http.StatusConflict})
	}

	newAppointment := models.Appointment{
		PatientID: patientID,
		DateTime:  appointmentTime,
	}

	if result := database.DB.Create(&newAppointment); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Não foi possível agendar a consulta", "code": http.StatusInternalServerError})
	}

	return c.JSON(http.StatusCreated, newAppointment)
}

// ListAppointments retorna todas as consultas do paciente autenticado.
func ListAppointments(c echo.Context) error {
	patientID, err := getPatientIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Acesso não autorizado", "code": http.StatusUnauthorized})
	}

	var appointments []models.Appointment
	database.DB.Where("patient_id = ?", patientID).Order("date_time asc").Find(&appointments)

	return c.JSON(http.StatusOK, appointments)
}

// CancelAppointment remove uma consulta específica.
func CancelAppointment(c echo.Context) error {
	appointmentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "ID da consulta inválido", "code": http.StatusBadRequest})
	}

	patientID, err := getPatientIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Acesso não autorizado", "code": http.StatusUnauthorized})
	}

	var appointment models.Appointment
	// Procura a consulta garantindo que o ID da consulta E o ID do paciente correspondem.
	// Isso impede que um paciente cancele a consulta de outro.
	if result := database.DB.Where("id = ? AND patient_id = ?", appointmentID, patientID).First(&appointment); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusForbidden, map[string]interface{}{"message": "Consulta não encontrada ou acesso não autorizado", "code": http.StatusForbidden})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Erro no servidor", "code": http.StatusInternalServerError})
	}

	// Deleta a consulta encontrada.
	if result := database.DB.Delete(&appointment); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Falha ao cancelar a consulta", "code": http.StatusInternalServerError})
	}

	// Retorna uma resposta vazia com status 204, indicando sucesso na exclusão.
	return c.NoContent(http.StatusNoContent)
}
