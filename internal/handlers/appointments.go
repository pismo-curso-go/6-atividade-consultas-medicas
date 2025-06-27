package handlers

import (
	"database/sql"
	"net/http"
	"saudemais-api/internal/models"
	"time"

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
