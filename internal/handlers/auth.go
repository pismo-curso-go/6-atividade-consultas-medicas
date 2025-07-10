package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"saudemais-api/internal/models"

	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Nome  string `json:"name"`
	Email string `json:"email"`
	Senha string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

		if req.Nome == "" || req.Email == "" || req.Senha == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Nome, email e senha são obrigatórios",
				"code":    400,
			})
		}

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

func Login(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req LoginRequest

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Dados inválidos",
				"code":    400,
			})
		}

		if req.Email == "" || req.Password == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Email e senha são obrigatórios",
				"code":    400,
			})
		}

		paciente, err := models.BuscarPacientePorEmail(db, req.Email)
		if err != nil {
			if err.Error() == "Paciente não encontrado" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": err.Error(),
					"code":    401,
				})
			}

			log.Println("Erro ao buscar paciente:", err)

			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Erro ao processar login",
				"code":    500,
			})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(paciente.PasswordHash), []byte(req.Password)); err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Credenciais inválidas",
				"code":    401,
			})
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "segredo_padrao"
		}

		claims := &jwt.RegisteredClaims{
			Subject:   req.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(secret))
		if err != nil {
			log.Println("Erro ao gerar token:", err)

			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Erro ao gerar token",
				"code":    500,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"token": t,
		})
	}
}
