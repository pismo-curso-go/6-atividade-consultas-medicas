package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"saudemais-api/internal/models"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Nome  string `json: "name"`
	Email string `json: "email"`
	Senha string `json: "password"`
}

func Register(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req RegisterRequest

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Dados inválidos",
				"code":    400,
			})
		}

		// if req.Nome == "" || req.Email == "" || req.Senha == "" {
		// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		// 		"message": "Nome, email e senha são obrigatórios",
		// 		"code":    400,
		// 	})
		// }

		err := models.CriarPaciente(db, req.Nome, req.Email, req.Senha)
		if err != nil {
			if err.Error() == "Paciente já cadastrado" {
				return c.JSON(http.StatusConflict, map[string]interface{}{
					"message": err.Error(),
					"code":    409,
				})
			}

			log.Println("Erro ao registrar paciente:", err)

			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Erro ao registrar paciente",
				"code":    500,
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Paciente registrado com sucesso",
			"code":    201,
		})
	}
}
