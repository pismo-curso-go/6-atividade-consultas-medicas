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
		var input struct {
			Email string `json:"email"`
			Senha string `json:"senha"`
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"erro": "Dados inválidos"})
		}

		var id int
		var passwordHash string
		err := db.QueryRow("SELECT id, password_hash FROM pacientes WHERE email = $1", input.Email).Scan(&id, &passwordHash)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusUnauthorized, map[string]string{"erro": "Usuário não encontrado"})
		} else if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"erro": "Erro ao buscar usuário"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(input.Senha)); err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"erro": "Senha incorreta"})
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "segredo_padrao"
		}

		claims := &jwt.RegisteredClaims{
			Subject:   input.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(secret))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"erro": "Erro ao gerar token"})
		}

		return c.JSON(http.StatusOK, map[string]string{"token": t})
	}
}
