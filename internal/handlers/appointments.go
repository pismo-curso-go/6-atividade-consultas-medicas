package handlers

import (
	"database/sql"
	"net/http"
	"saudemais-api/internal/models"
	"time"
	"strconv"
	"github.com/labstack/echo/v4"
)

type AgendarRequest struct {
	Datetime string `json:"datetime"`
}


func AgendarConsulta(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req AgendarRequest

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Dados inválidos",
				"code":    400,
			})
		}

		if req.Datetime == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Data e hora são obrigatórios",
				"code":    400,
			})
		}

		agendamento, err := time.Parse("2006-01-02T15:04:05", req.Datetime)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Formato de data inválido",
				"code":    400,
			})
		}

		if agendamento.Before(time.Now()) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Consulta no passado não é permitida",
				"code":    400,
			})
		}

		email := c.Get("userEmail").(string)

		paciente, err := models.BuscarPacientePorEmail(db, email)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Paciente não encontrado",
				"code":    401,
			})
		}

		var consultaID int
		err = db.QueryRow("SELECT id FROM consultas WHERE patient_id = $1 AND datetime = $2", paciente.ID, agendamento).
			Scan(&consultaID)

		if err != sql.ErrNoRows && err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Erro ao verificar agendamento",
				"code":    500,
			})
		}

		if consultaID != 0 {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"message": "Consulta já agendada para este horário",
				"code":    409,
			})
		}

		_, err = db.Exec("INSERT INTO consultas (patient_id, datetime) VALUES ($1, $2)", paciente.ID, agendamento)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Erro ao agendar consulta",
				"code":    500,
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Consulta agendada com sucesso",
			"code":    201,
		})
	}

}

func ListarConsulta(db *sql.DB) echo.HandlerFunc {

	return func(c echo.Context) error {
		email, ok := c.Get("userEmail").(string)
		if !ok || email == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Usuário não autenticado",
				"code":    401,
			})
		}

		paciente, err := models.BuscarPacientePorEmail(db, email)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Paciente não encontrado",
				"code":    401,
			})
		}

		consultas, err := models.ListaConsultasPorPaciente(db, paciente.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Erro ao listar consultas",
				"code":    500,
			})
		}

		// Define a struct de resposta formatada
		type ConsultaResponse struct {
			ID       int    `json:"id"`
			Datetime string `json:"datetime"`
		}

		var resposta []ConsultaResponse
		for _, consulta := range consultas {
			resposta = append(resposta, ConsultaResponse{
				ID:       consulta.ID,
				Datetime: consulta.Datetime.Format("2006-01-02 15:04:05"),
			})
		}

		return c.JSON(http.StatusOK, resposta)
	}
}

func CancelarConsulta(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Pega o email do paciente autenticado
		email, ok := c.Get("userEmail").(string)
		if !ok || email == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Usuário não autenticado",
				"code":    401,
			})
		}

		// Busca o paciente pelo e-mail
		paciente, err := models.BuscarPacientePorEmail(db, email)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Paciente não encontrado",
				"code":    401,
			})
		}

		// Converte o ID da consulta (rota: /appointments/:id)
		idStr := c.Param("id")
		consultaID, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "ID inválido",
				"code":    400,
			})
		}

		// Cancela a consulta no banco de dados
		result, err := db.Exec("DELETE FROM consultas WHERE id = $1 AND patient_id = $2", consultaID, paciente.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Erro ao cancelar consulta",
				"code":    500,
			})
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Consulta não encontrada ou não pertence ao paciente",
				"code":    404,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Consulta cancelada com sucesso",
			"code":    200,
		})
	}
}
